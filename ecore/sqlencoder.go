package ecore

import (
	"database/sql"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

type sqlEncoderFeatureData struct {
	table       string
	factory     EFactory
	dataType    EDataType
	featureKind sqlFeatureKind
	isTransient bool
}

type sqlEncoderObjectData struct {
	id        int64
	classData *sqlEncoderClassData
}

type sqlEncoderClassData struct {
	id          int64
	table       string
	featureData []*sqlEncoderFeatureData
}

type sqlEncoderPackageData struct {
	id int64
}

type SQLEncoder struct {
	resource          EResource
	writer            io.Writer
	driver            string
	db                *sql.DB
	packageDataMap    map[EPackage]*sqlEncoderPackageData
	classDataMap      map[EClass]*sqlEncoderClassData
	insertContentStmt *sql.Stmt
	insertPackageStmt *sql.Stmt
	insertClassStmt   *sql.Stmt
	insertObjectStmts map[EClass]*sql.Stmt
	idAttributeName   string
}

func NewSQLEncoder(resource EResource, w io.Writer, options map[string]any) *SQLEncoder {
	e := &SQLEncoder{
		resource:          resource,
		writer:            w,
		driver:            "sqlite",
		packageDataMap:    map[EPackage]*sqlEncoderPackageData{},
		classDataMap:      map[EClass]*sqlEncoderClassData{},
		insertObjectStmts: map[EClass]*sql.Stmt{},
	}
	if options != nil {
		if driver, isDriver := options[SQL_OPTION_DRIVER]; isDriver {
			e.driver = driver.(string)
		}
		e.idAttributeName, _ = options[JSON_OPTION_ID_ATTRIBUTE_NAME].(string)
	}
	return e
}

func (e *SQLEncoder) createDB() (*sql.DB, error) {
	fileName := filepath.Base(e.resource.GetURI().Path())
	dbPath, err := sqlTmpDB(fileName)
	if err != nil {
		return nil, err
	}

	// open db
	db, err := sql.Open(e.driver, dbPath)
	if err != nil {
		return nil, err
	}

	// version info
	version := fmt.Sprintf(`PRAGMA user_version = %v`, sqlCodecVersion)
	_, err = db.Exec(version)
	if err != nil {
		return nil, err
	}

	// properties infos
	properties := `
	PRAGMA synchronous = NORMAL;
	PRAGMA journal_mode = WAL;
	`
	_, err = db.Exec(properties)
	if err != nil {
		return nil, err
	}

	// common tables
	tables := `
	CREATE TABLE packages ( 
		packageID INTEGER PRIMARY KEY AUTOINCREMENT,
		uri TEXT
	);
	CREATE TABLE classes (
		classID INTEGER PRIMARY KEY AUTOINCREMENT,
		packageID INTEGER,
		name TEXT,
		FOREIGN KEY(packageID) REFERENCES packages(packageID)
	);
	CREATE TABLE contents (
		classID INTEGER,
		objectID INTEGER,
		FOREIGN KEY(classID) REFERENCES classes(classID)
	);
	`
	_, err = db.Exec(tables)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (e *SQLEncoder) EncodeResource() {
	var err error
	e.db, err = e.createDB()
	if err != nil {
		e.resource.GetErrors().Add(NewEDiagnosticImpl(err.Error(), e.resource.GetURI().String(), 0, 0))
		return
	}
	defer func() {
		_ = e.db.Close()
	}()

	if contents := e.resource.GetContents(); !contents.Empty() {
		object := contents.Get(0).(EObject)
		if err := e.encodeContent(object); err != nil {
			e.resource.GetErrors().Add(NewEDiagnosticImpl(err.Error(), e.resource.GetURI().String(), 0, 0))
			return
		}
	}
}

func (e *SQLEncoder) EncodeObject(object EObject) error {
	return nil
}

func (e *SQLEncoder) encodeContent(eObject EObject) error {
	objectData, err := e.encodeObject(eObject)
	if err != nil {
		return err
	}

	insertContentStmt, err := e.getInsertContentStmt()
	if err != nil {
		return err
	}

	if _, err := insertContentStmt.Exec(objectData.classData.id, objectData.id); err != nil {
		return err
	}

	return nil
}

func (e *SQLEncoder) encodeObject(eObject EObject) (*sqlEncoderObjectData, error) {
	// encode object class
	eClass := eObject.EClass()
	classData, err := e.encodeClass(eClass)
	if err != nil {
		return nil, err
	}

	idManager := e.resource.GetObjectIDManager()
	if len(e.idAttributeName) == 0 {
		idManager = nil
	}

	// create table
	insertObjectStmt, err := e.getInsertObjectStmt(eClass, classData, idManager)
	if err != nil {
		return nil, err
	}

	// insert object in table
	args := []any{}
	if idManager != nil {
		args = append(args, idManager.GetID(eObject))
	}

	sqlResult, err := insertObjectStmt.Exec(args...)
	if err != nil {
		return nil, err
	}

	// retrieve object id
	objectID, err := sqlResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	// encode features values
	for featureID, featureData := range classData.featureData {
		e.encodeFeatureValue(eObject, featureID, objectID, featureData)
	}

	return &sqlEncoderObjectData{
		id:        objectID,
		classData: classData,
	}, nil
}

func (e *SQLEncoder) encodeFeatureValue(eObject EObject, featureID int, objectID int64, featureData *sqlEncoderFeatureData) {

}

func (e *SQLEncoder) encodeClass(eClass EClass) (*sqlEncoderClassData, error) {
	classData := e.classDataMap[eClass]
	if classData == nil {
		// encode package
		ePackage := eClass.GetEPackage()
		packageData, err := e.encodePackage(ePackage)
		if err != nil {
			return nil, err
		}

		// insert class in sql
		insertClassStmt, err := e.getInsertClassStmt()
		if err != nil {
			return nil, err
		}
		sqlResult, err := insertClassStmt.Exec(packageData.id, eClass.GetName())
		if err != nil {
			return nil, err
		}

		// retrieve class index
		id, err := sqlResult.LastInsertId()
		if err != nil {
			return nil, err
		}

		// create class data
		classData = e.newClassData(eClass, id)
		e.classDataMap[eClass] = classData
	}
	return classData, nil
}

func (e *SQLEncoder) encodePackage(ePackage EPackage) (*sqlEncoderPackageData, error) {
	ePackageData := e.packageDataMap[ePackage]
	if ePackageData == nil {
		// insert new package
		insertPackageStmt, err := e.getInsertPackageStmt()
		if err != nil {
			return nil, err
		}
		sqlResult, err := insertPackageStmt.Exec(ePackage.GetNsURI())
		if err != nil {
			return nil, err
		}
		// retrieve package index
		id, err := sqlResult.LastInsertId()
		if err != nil {
			return nil, err
		}
		// create data
		ePackageData = e.newPackageData(id)
		e.packageDataMap[ePackage] = ePackageData
	}
	return ePackageData, nil
}

func (e *SQLEncoder) getInsertContentStmt() (*sql.Stmt, error) {
	if e.insertContentStmt == nil {
		insertContentStmt, err := e.db.Prepare(`INSERT INTO contents (classID,objectID) VALUES (?,?)`)
		if err != nil {
			return nil, err
		}
		e.insertContentStmt = insertContentStmt
	}
	return e.insertContentStmt, nil
}

func (e *SQLEncoder) getInsertObjectStmt(eClass EClass, classData *sqlEncoderClassData, idManager EObjectIDManager) (*sql.Stmt, error) {
	insertObjectStmt, isObjectInsertStmt := e.insertObjectStmts[eClass]
	if !isObjectInsertStmt {
		var err error
		// create table
		{
			var query strings.Builder
			query.WriteString("CREATE TABLE ")
			query.WriteString(classData.table)
			query.WriteString(" (objectID INTEGER PRIMARY KEY AUTOINCREMENT")
			if idManager != nil {
				query.WriteString(",")
				query.WriteString(e.idAttributeName)
			}
			query.WriteString(")")

			if _, err = e.db.Exec(query.String()); err != nil {
				return nil, err
			}
		}
		// create stmt
		{
			var query strings.Builder
			query.WriteString("INSERT INTO ")
			query.WriteString(classData.table)
			query.WriteString(" (objectID")

			if idManager != nil {
				query.WriteString(",")
				query.WriteString(e.idAttributeName)
				query.WriteString(") VALUES (NULL,?)")
			} else {
				query.WriteString(") VALUES (NULL)")
			}

			insertObjectStmt, err = e.db.Prepare(query.String())
			if err != nil {
				return nil, err
			}
		}
		// set stmt
		e.insertObjectStmts[eClass] = insertObjectStmt
	}
	return insertObjectStmt, nil
}

func (e *SQLEncoder) getInsertClassStmt() (*sql.Stmt, error) {
	if e.insertClassStmt == nil {
		stmt, err := e.db.Prepare(`INSERT INTO classes (packageID,name) VALUES (?,?)`)
		if err != nil {
			return nil, err
		}
		e.insertClassStmt = stmt
	}
	return e.insertClassStmt, nil
}

func (e *SQLEncoder) getInsertPackageStmt() (*sql.Stmt, error) {
	// create statement if needed
	if e.insertPackageStmt == nil {
		stmt, err := e.db.Prepare(`INSERT INTO packages (uri) VALUES (?)`)
		if err != nil {
			return nil, err
		}
		e.insertPackageStmt = stmt
	}
	return e.insertPackageStmt, nil
}

func (e *SQLEncoder) newPackageData(id int64) *sqlEncoderPackageData {
	return &sqlEncoderPackageData{
		id: id,
	}
}

func (e *SQLEncoder) newClassData(eClass EClass, id int64) *sqlEncoderClassData {
	// create data
	ePackage := eClass.GetEPackage()
	eFeatures := eClass.GetEAllStructuralFeatures()
	classData := &sqlEncoderClassData{
		id:          id,
		table:       ePackage.GetNsPrefix() + "_" + strings.ToLower(eClass.GetName()),
		featureData: make([]*sqlEncoderFeatureData, 0, eFeatures.Size()),
	}
	// class features data
	for itFeature := eFeatures.Iterator(); itFeature.HasNext(); {
		eFeature := itFeature.Next().(EStructuralFeature)
		classData.featureData = append(classData.featureData, e.newFeatureData(eFeature, classData.table))
	}
	return classData
}

func (e *SQLEncoder) newFeatureData(eFeature EStructuralFeature, classTableName string) *sqlEncoderFeatureData {
	featureData := &sqlEncoderFeatureData{
		table:       classTableName + "_" + strings.ToLower(eFeature.GetName()),
		featureKind: getSQLCodecFeatureKind(eFeature),
	}
	if eReference, _ := eFeature.(EReference); eReference != nil {
		featureData.isTransient = eReference.IsTransient() || (eReference.IsContainer() && !eReference.IsResolveProxies())
	} else if eAttribute, _ := eFeature.(EAttribute); eAttribute != nil {
		eDataType := eAttribute.GetEAttributeType()
		featureData.isTransient = eAttribute.IsTransient()
		featureData.dataType = eDataType
		featureData.factory = eDataType.GetEPackage().GetEFactoryInstance()
	}
	return featureData
}
