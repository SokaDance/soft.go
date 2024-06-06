package ecore

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type sqlObjectManager interface {
	registerObject(EObject)
}

type sqlDecoderIDManager interface {
	sqlCodecIDManager
	getPackageFromID(int64) (EPackage, bool)
	getObjectFromID(int64) (EObject, bool)
	getClassFromID(int64) (EClass, bool)
	getEnumLiteralFromID(int64) (EEnumLiteral, bool)
}

type sqlDecoderClassData struct {
	eClass   EClass
	eFactory EFactory
}

type sqlDecoderIDManagerImpl struct {
	packages     map[int64]EPackage
	objects      map[int64]EObject
	classes      map[int64]EClass
	enumLiterals map[int64]EEnumLiteral
}

func newSqlDecoderIDManager() *sqlDecoderIDManagerImpl {
	return &sqlDecoderIDManagerImpl{
		packages:     map[int64]EPackage{},
		objects:      map[int64]EObject{},
		classes:      map[int64]EClass{},
		enumLiterals: map[int64]EEnumLiteral{},
	}
}

func (r *sqlDecoderIDManagerImpl) setPackageID(p EPackage, id int64) {
	r.packages[id] = p
}

func (r *sqlDecoderIDManagerImpl) getPackageFromID(id int64) (p EPackage, b bool) {
	p, b = r.packages[id]
	return
}

func (r *sqlDecoderIDManagerImpl) setObjectID(o EObject, id int64) {
	r.objects[id] = o

	// set sql id if created object is an sql object
	if sqlObject, _ := o.(SQLObject); sqlObject != nil {
		sqlObject.SetSqlID(id)
	}
}

func (r *sqlDecoderIDManagerImpl) getObjectFromID(id int64) (o EObject, b bool) {
	o, b = r.objects[id]
	return
}

func (r *sqlDecoderIDManagerImpl) setClassID(c EClass, id int64) {
	r.classes[id] = c
}

func (r *sqlDecoderIDManagerImpl) getClassFromID(id int64) (c EClass, b bool) {
	c, b = r.classes[id]
	return
}

func (r *sqlDecoderIDManagerImpl) setEnumLiteralID(e EEnumLiteral, id int64) {
	r.enumLiterals[id] = e
}

func (r *sqlDecoderIDManagerImpl) getEnumLiteralFromID(id int64) (e EEnumLiteral, b bool) {
	e, b = r.enumLiterals[id]
	return
}

type sqlDecoderObjectManager struct {
}

func newSqlDecoderObjectManager() *sqlDecoderObjectManager {
	return &sqlDecoderObjectManager{}
}

func (r *sqlDecoderObjectManager) registerObject(EObject) {
}

type sqlDecoder struct {
	*sqlBase
	classDataMap     map[EClass]*sqlDecoderClassData
	packageRegistry  EPackageRegistry
	sqlIDManager     sqlDecoderIDManager
	sqlObjectManager sqlObjectManager
}

func (d *sqlDecoder) resolveURI(uri *URI) *URI {
	if d.uri != nil {
		return d.uri.Resolve(uri)
	}
	return uri
}

func (d *sqlDecoder) decodeAny(stmt *sqlite.Stmt, i int) any {
	switch stmt.ColumnType(i) {
	case sqlite.TypeNull:
		return nil
	case sqlite.TypeInteger:
		return stmt.ColumnInt64(i)
	case sqlite.TypeText:
		return stmt.ColumnText(i)
	case sqlite.TypeFloat:
		return stmt.ColumnFloat(i)
	case sqlite.TypeBlob:
		bytes := make([]byte, stmt.ColumnLen(i))
		stmt.ColumnBytes(i, bytes)
		return bytes
	default:
		panic("sqlite type not supported")
	}
}

func (d *sqlDecoder) decodePackage(conn *sqlite.Conn, id int64) (EPackage, error) {
	ePackage, isPackage := d.sqlIDManager.getPackageFromID(id)
	if !isPackage {
		table := d.schema.packagesTable
		var packageID int64
		var packageURI string
		if err := sqlitex.Execute(
			conn,
			table.selectQuery(nil, table.keyName()+"=?", ""),
			&sqlitex.ExecOptions{
				Args: []any{id},
				ResultFunc: func(stmt *sqlite.Stmt) error {
					packageID = stmt.ColumnInt64(0)
					packageURI = stmt.ColumnText(1)
					return nil
				},
			}); err != nil {
			return nil, err
		}

		// retrieve package
		if d.packageRegistry == nil {
			panic(fmt.Errorf("package registry not defined in sql decoder"))
		}
		ePackage = d.packageRegistry.GetPackage(packageURI)
		if ePackage == nil {
			return nil, fmt.Errorf("unable to find package '%s'", packageURI)
		}

		// register package id
		d.sqlIDManager.setPackageID(ePackage, packageID)
	}
	return ePackage, nil
}

func (d *sqlDecoder) decodeClass(conn *sqlite.Conn, id int64) (*sqlDecoderClassData, error) {
	eClass, isClass := d.sqlIDManager.getClassFromID(id)
	if !isClass {
		table := d.schema.classesTable
		var className string
		var packageID int64
		if err := sqlitex.Execute(
			conn,
			table.selectQuery(nil, table.keyName()+"=?", ""),
			&sqlitex.ExecOptions{
				Args: []any{id},
				ResultFunc: func(stmt *sqlite.Stmt) error {
					packageID = stmt.ColumnInt64(1)
					className = stmt.ColumnText(2)
					return nil
				},
			}); err != nil {
			return nil, err
		}

		// retrieve package
		ePackage, err := d.decodePackage(conn, packageID)
		if err != nil {
			return nil, err
		}
		// retrieve class
		eClass, _ = ePackage.GetEClassifier(className).(EClass)
		if eClass == nil {
			return nil, fmt.Errorf("unable to find class '%s' in package '%s'", className, ePackage.GetNsURI())
		}
		// set class id
		d.sqlIDManager.setClassID(eClass, id)
	}
	return d.getDecoderClassData(eClass), nil

}

func (d *sqlDecoder) getDecoderClassData(eClass EClass) *sqlDecoderClassData {
	classData, isClassData := d.classDataMap[eClass]
	if !isClassData {
		classData = &sqlDecoderClassData{
			eClass:   eClass,
			eFactory: eClass.GetEPackage().GetEFactoryInstance(),
		}
		d.classDataMap[eClass] = classData
	}
	return classData
}

func (d *sqlDecoder) decodeContents(conn *sqlite.Conn) ([]EObject, error) {
	table := d.schema.contentsTable
	contents := []EObject{}
	if err := sqlitex.Execute(
		conn,
		table.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				// retrieve object id
				objectID := stmt.ColumnInt64(0)
				// decode object
				object, err := d.decodeObject(conn, objectID)
				if err != nil {
					return err
				}
				// add object to contents
				contents = append(contents, object)
				return nil
			},
		}); err != nil {
		return nil, err
	}
	return contents, nil
}

func (d *sqlDecoder) decodeObject(conn *sqlite.Conn, id int64) (EObject, error) {
	eObject, isObject := d.sqlIDManager.getObjectFromID(id)
	if !isObject {
		table := d.schema.objectsTable
		var classID int64
		var objectID string
		var isObjectID bool
		if err := sqlitex.Execute(
			conn,
			table.selectQuery(nil, table.keyName()+"=?", ""),
			&sqlitex.ExecOptions{
				Args: []any{id},
				ResultFunc: func(stmt *sqlite.Stmt) error {
					classID = stmt.ColumnInt64(1)
					isObjectID = stmt.ColumnCount() > 2
					if isObjectID {
						objectID = stmt.ColumnText(2)
					}
					return nil
				},
			}); err != nil {
			return nil, err
		}

		// retrieve class data
		classData, err := d.decodeClass(conn, classID)
		if err != nil {
			return nil, err
		}

		// create object
		eObject = classData.eFactory.Create(classData.eClass)

		// register its id
		if isObjectID && d.idManager != nil {
			if err := d.idManager.SetID(eObject, objectID); err != nil {
				return nil, err
			}
		}

		// register in sql id manager
		d.sqlIDManager.setObjectID(eObject, id)

		// register in sql object maneger
		d.sqlObjectManager.registerObject(eObject)
	}
	return eObject, nil
}

func (d *sqlDecoder) decodeEnumLiteral(conn *sqlite.Conn, id int64) (EEnumLiteral, error) {
	eEnumLiteral, isEnumLiteral := d.sqlIDManager.getEnumLiteralFromID(id)
	if !isEnumLiteral {
		table := d.schema.enumsTable
		var enumID int64
		var packageID int64
		var enumName string
		var literalValue string
		if err := sqlitex.Execute(
			conn,
			table.selectQuery(nil, table.keyName()+"=?", ""),
			&sqlitex.ExecOptions{
				Args: []any{id},
				ResultFunc: func(stmt *sqlite.Stmt) error {
					enumID = stmt.ColumnInt64(0)
					packageID = stmt.ColumnInt64(1)
					enumName = stmt.ColumnText(2)
					literalValue = stmt.ColumnText(3)
					return nil
				},
			}); err != nil {
			return nil, err
		}

		// package
		ePackage, err := d.decodePackage(conn, packageID)
		if err != nil {
			return nil, err
		}

		// enum
		eEnum, _ := ePackage.GetEClassifier(enumName).(EEnum)
		if eEnum == nil {
			return nil, fmt.Errorf("unable to find enum '%s' in package '%s'", enumName, ePackage.GetName())
		}

		eEnumLiteral = eEnum.GetEEnumLiteralByLiteral(literalValue)
		if eEnumLiteral == nil {
			return nil, fmt.Errorf("unable to find enum literal '%s' in enum '%s'", literalValue, eEnum.GetName())
		}

		// register enum value
		d.sqlIDManager.setEnumLiteralID(eEnumLiteral, enumID)
	}
	return eEnumLiteral, nil
}

func (d *sqlDecoder) decodeFeatureValue(conn *sqlite.Conn, featureData *sqlFeatureSchema, value any) (any, error) {
	switch featureData.featureKind {
	case sfkObject, sfkObjectList:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case int64:
			return d.decodeObject(conn, v)
		default:
			return nil, fmt.Errorf("%v is not supported as a object id", v)
		}
	case sfkObjectReference, sfkObjectReferenceList:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case string:
			// no reference
			if len(v) == 0 {
				return nil, nil
			} else {
				sqlID, err := strconv.ParseInt(v, 10, 64)
				if err == nil {
					// string is an int => object is in the resource
					// decode it
					return d.decodeObject(conn, sqlID)
				} else {
					// string is an uri
					proxyURI := NewURI(v)
					resolvedURI := d.resolveURI(proxyURI)
					// create proxy
					eFeature := featureData.feature
					eClass := eFeature.GetEType().(EClass)
					eFactory := eClass.GetEPackage().GetEFactoryInstance()
					eObject := eFactory.Create(eClass)
					eObjectInternal := eObject.(EObjectInternal)
					eObjectInternal.ESetProxyURI(resolvedURI)
					return eObject, nil
				}
			}
		default:
			return nil, fmt.Errorf("%v is not supported as a object reference uri", v)
		}
	case sfkBool:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case bool:
			return v, nil
		default:
			return nil, fmt.Errorf("%v is not a bool value", v)
		}
	case sfkByte:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case byte:
			return v, nil
		default:
			return nil, fmt.Errorf("%v is not a bool value", v)
		}
	case sfkInt:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case int64:
			return int(v), nil
		default:
			return nil, fmt.Errorf("%v is not a int value", v)
		}
	case sfkInt64:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case int64:
			return v, nil
		default:
			return nil, fmt.Errorf("%v is not a int64 value", v)
		}
	case sfkInt32:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case int64:
			return int32(v), nil
		default:
			return nil, fmt.Errorf("%v is not a int32 value", v)
		}
	case sfkInt16:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case int64:
			return int16(v), nil
		default:
			return nil, fmt.Errorf("%v is not a int16 value", v)
		}
	case sfkEnum:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case int64:
			enumLiteral, err := d.decodeEnumLiteral(conn, v)
			if err != nil {
				return nil, err
			}
			instance := enumLiteral.GetInstance()
			if instance == nil {
				instance = d.decodeFeatureData(featureData, enumLiteral.GetLiteral())
			}
			return instance, nil
		default:
			return nil, fmt.Errorf("%v is not a enum value", value)
		}
	case sfkString:
		switch v := value.(type) {
		case nil:
			return "", nil
		case string:
			return v, nil
		default:
			return "", fmt.Errorf("%v is not a string value", v)
		}
	case sfkByteArray:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case []byte:
			return v, nil
		default:
			return "", fmt.Errorf("%v is not a byte array value", v)
		}
	case sfkDate:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case time.Time:
			return &v, nil
		case string:
			t, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return nil, err
			}
			return &t, nil
		default:
			return nil, fmt.Errorf("%v is not a time value", v)
		}
	case sfkFloat64:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case float64:
			return v, nil
		default:
			return nil, fmt.Errorf("%v is not a float64 value", value)
		}
	case sfkFloat32:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case float64:
			return float32(v), nil
		default:
			return nil, fmt.Errorf("%v is not a float64 value", value)
		}
	case sfkData, sfkDataList:
		switch v := value.(type) {
		case nil:
			return nil, nil
		case []byte:
			return d.decodeFeatureData(featureData, string(v)), nil
		case string:
			return d.decodeFeatureData(featureData, v), nil
		default:
			return nil, fmt.Errorf("%v is not a data value", value)
		}
	}

	return nil, nil
}

func (d *sqlDecoder) decodeFeatureData(featureSchema *sqlFeatureSchema, v string) any {
	eFeature := featureSchema.feature
	eDataType := eFeature.GetEType().(EDataType)
	eFactory := eDataType.GetEPackage().GetEFactoryInstance()
	return eFactory.CreateFromString(eDataType, v)
}

type SQLDecoder struct {
	sqlDecoder
	resource     EResource
	driver       string
	connProvider func() (*sqlite.Conn, error)
	connClose    func(conn *sqlite.Conn) error
}

func NewSQLReaderDecoder(r io.Reader, resource EResource, options map[string]any) *SQLDecoder {
	return newSQLDecoder(
		func() (*sqlite.Conn, error) {
			fileName := filepath.Base(resource.GetURI().Path())
			dbPath, err := sqlTmpDB(fileName)
			if err != nil {
				return nil, err
			}

			dbFile, err := os.Create(dbPath)
			if err != nil {
				return nil, err
			}

			_, err = io.Copy(dbFile, r)
			if err != nil {
				dbFile.Close()
				return nil, err
			}
			dbFile.Close()

			return sqlite.OpenConn(dbPath, sqlite.OpenReadOnly, sqlite.OpenWAL)

		},
		func(conn *sqlite.Conn) error {
			return conn.Close()
		},
		resource,
		options)
}

func NewSQLDBDecoder(conn *sqlite.Conn, resource EResource, options map[string]any) *SQLDecoder {
	return newSQLDecoder(
		func() (*sqlite.Conn, error) {
			return conn, nil
		},
		func(db *sqlite.Conn) error {
			return nil
		},
		resource,
		options,
	)
}

func newSQLDecoder(connProvider func() (*sqlite.Conn, error), connClose func(conn *sqlite.Conn) error, resource EResource, options map[string]any) *SQLDecoder {
	// options
	schemaOptions := []sqlSchemaOption{}
	driver := "sqlite"
	idAttributeName := ""
	if options != nil {
		if d, isDriver := options[SQL_OPTION_DRIVER]; isDriver {
			driver = d.(string)
		}

		idAttributeName, _ = options[SQL_OPTION_ID_ATTRIBUTE_NAME].(string)
		if resource.GetObjectIDManager() != nil && len(idAttributeName) > 0 {
			schemaOptions = append(schemaOptions, withIDAttributeName(idAttributeName))
		}
	}

	// package registry
	packageRegistry := GetPackageRegistry()
	resourceSet := resource.GetResourceSet()
	if resourceSet != nil {
		packageRegistry = resourceSet.GetPackageRegistry()
	}

	return &SQLDecoder{
		sqlDecoder: sqlDecoder{
			sqlBase: &sqlBase{
				uri:             resource.GetURI(),
				idManager:       resource.GetObjectIDManager(),
				idAttributeName: idAttributeName,
				schema:          newSqlSchema(schemaOptions...),
			},
			packageRegistry:  packageRegistry,
			sqlIDManager:     newSqlDecoderIDManager(),
			sqlObjectManager: newSqlDecoderObjectManager(),
			classDataMap:     map[EClass]*sqlDecoderClassData{},
		},
		resource:     resource,
		connProvider: connProvider,
		connClose:    connClose,
		driver:       driver,
	}
}

func (d *SQLDecoder) DecodeResource() {
	conn, err := d.connProvider()
	if err != nil {
		d.addError(err)
		return
	}
	defer func() {
		if err := d.connClose(conn); err != nil {
			d.addError(err)
		}
	}()
	if err := d.decodeVersion(conn); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodePackages(conn); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeEnums(conn); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeClasses(conn); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeObjects(conn); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeFeatures(conn); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeContents(conn); err != nil {
		d.addError(err)
		return
	}
}

func (d *SQLDecoder) DecodeObject() (EObject, error) {
	panic("SQLDecoder doesn't support object decoding")
}

func (d *SQLDecoder) decodeVersion(conn *sqlite.Conn) error {
	var version int64
	if err := sqlitex.Execute(conn, "PRAGMA user_version;", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			version = stmt.ColumnInt64(0)
			return nil
		},
	}); err != nil {
		return err
	}

	if version != sqlCodecVersion {
		return fmt.Errorf("history version %v is not supported", version)
	}
	return nil
}

func (d *SQLDecoder) decodeContents(conn *sqlite.Conn) error {
	return sqlitex.Execute(
		conn,
		d.schema.contentsTable.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				objectID := stmt.ColumnInt64(0)
				object, _ := d.sqlIDManager.getObjectFromID(objectID)
				if object == nil {
					return fmt.Errorf("unable to find object with id '%v'", objectID)
				}
				d.resource.GetContents().Add(object)
				return nil
			},
		})
}

func (d *SQLDecoder) decodePackages(conn *sqlite.Conn) error {
	return sqlitex.Execute(
		conn,
		d.schema.packagesTable.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				packageID := stmt.ColumnInt64(0)
				packageURI := stmt.ColumnText(1)
				ePackage := d.packageRegistry.GetPackage(packageURI)
				if ePackage == nil {
					return fmt.Errorf("unable to find package '%s'", packageURI)
				}
				d.sqlIDManager.setPackageID(ePackage, packageID)
				return nil
			},
		})
}

func (d *SQLDecoder) decodeEnums(conn *sqlite.Conn) error {
	return sqlitex.Execute(
		conn,
		d.schema.enumsTable.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				enumID := stmt.ColumnInt64(0)
				packageID := stmt.ColumnInt64(1)
				enumName := stmt.ColumnText(2)
				literalValue := stmt.ColumnText(3)

				ePackage, _ := d.sqlIDManager.getPackageFromID(packageID)
				if ePackage == nil {
					return fmt.Errorf("unable to find package with id '%d'", packageID)
				}

				eEnum, _ := ePackage.GetEClassifier(enumName).(EEnum)
				if eEnum == nil {
					return fmt.Errorf("unable to find enum '%s' in package '%s'", enumName, ePackage.GetNsURI())
				}

				eEnumLiteral := eEnum.GetEEnumLiteralByLiteral(literalValue)
				if eEnumLiteral == nil {
					return fmt.Errorf("unable to find enum literal '%s' in enum '%s'", literalValue, eEnum.GetName())
				}

				// create map enum
				d.sqlIDManager.setEnumLiteralID(eEnumLiteral, enumID)
				return nil
			},
		})
}

func (d *SQLDecoder) decodeClasses(conn *sqlite.Conn) error {
	return sqlitex.Execute(
		conn,
		d.schema.classesTable.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				classID := stmt.ColumnInt64(0)
				packageID := stmt.ColumnInt64(1)
				className := stmt.ColumnText(2)
				ePackage, _ := d.sqlIDManager.getPackageFromID(packageID)
				if ePackage == nil {
					return fmt.Errorf("unable to find package with id '%d'", packageID)
				}
				eClass, _ := ePackage.GetEClassifier(className).(EClass)
				if eClass == nil {
					return fmt.Errorf("unable to find class '%s' in package '%s'", className, ePackage.GetNsURI())
				}

				d.sqlIDManager.setClassID(eClass, classID)

				d.classDataMap[eClass] = &sqlDecoderClassData{
					eClass:   eClass,
					eFactory: ePackage.GetEFactoryInstance(),
				}
				return nil
			},
		})
}

func (d *SQLDecoder) decodeObjects(conn *sqlite.Conn) error {
	return sqlitex.Execute(
		conn,
		d.schema.objectsTable.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				objectID := stmt.ColumnInt64(0)
				classID := stmt.ColumnInt64(1)
				eClass, _ := d.sqlIDManager.getClassFromID(classID)
				if eClass == nil {
					return fmt.Errorf("unable to find class with id '%v'", classID)
				}

				// class data
				classData := d.classDataMap[eClass]
				if classData == nil {
					return fmt.Errorf("unable to find class data with id '%v'", classID)
				}

				// create & register object
				eObject := classData.eFactory.Create(classData.eClass)
				d.sqlIDManager.setObjectID(eObject, objectID)

				// set its id
				if stmt.ColumnCount() > 2 {
					switch stmt.ColumnType(2) {
					case sqlite.TypeNull:
					case sqlite.TypeText:
						if d.idManager != nil {
							id := stmt.ColumnText(2)
							if err := d.idManager.SetID(eObject, id); err != nil {
								return err
							}
						}
					}
				}
				return nil
			},
		})
}

func (d *SQLDecoder) decodeFeatures(conn *sqlite.Conn) error {
	decoded := map[EClass]struct{}{}
	// for each leaf class
	for _, classData := range d.classDataMap {
		eClass := classData.eClass
		itSuper := eClass.GetEAllSuperTypes().Iterator()
		for eClass != nil {
			// decode class features
			if _, idDecoded := decoded[eClass]; !idDecoded {
				decoded[eClass] = struct{}{}
				if err := d.decodeClassFeatures(conn, eClass); err != nil {
					return err
				}
			}

			// next super class
			if itSuper.HasNext() {
				eClass = itSuper.Next().(EClass)
			} else {
				eClass = nil
			}
		}
	}
	return nil
}

func (d *SQLDecoder) decodeClassFeatures(conn *sqlite.Conn, eClass EClass) error {
	classSchema := d.schema.getClassSchema(eClass)
	columnFeatures := []*sqlFeatureSchema{}
	for _, featureData := range classSchema.features {
		if featureData.column != nil {
			columnFeatures = append(columnFeatures, featureData)
		} else if featureData.table != nil {
			if err := d.decodeTableFeature(conn, featureData.table, featureData); err != nil {
				return err
			}
		}
	}

	return d.decodeColumnFeatures(conn, classSchema.table, columnFeatures)
}

func (d *SQLDecoder) decodeColumnFeatures(conn *sqlite.Conn, table *sqlTable, columnFeatures []*sqlFeatureSchema) error {
	if len(columnFeatures) == 0 {
		return nil
	}
	return sqlitex.Execute(
		conn,
		table.selectQuery(nil, "", ""),
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				objectID := stmt.ColumnInt64(0)
				eObject, _ := d.sqlIDManager.getObjectFromID(objectID)
				if eObject == nil {
					return fmt.Errorf("unable to find object with id '%v'", objectID)
				}

				// for each column
				for i, columnData := range columnFeatures {
					value := d.decodeAny(stmt, i+1)
					columnValue, err := d.decodeFeatureValue(conn, columnData, value)
					if err != nil {
						return err
					}
					// column value is nil, if column is not set
					// so use default value
					if columnValue != nil {
						eObject.ESet(columnData.feature, columnValue)
					}
				}
				return nil
			},
		})
}

func (d *SQLDecoder) decodeTableFeature(conn *sqlite.Conn, table *sqlTable, tableFeature *sqlFeatureSchema) error {
	column := sqlEscapeIdentifier(table.columns[len(table.columns)-1].columnName)
	query := table.selectQuery([]string{table.keyName(), column}, "", table.keyName()+" ASC, idx ASC")
	feature := tableFeature.feature
	values := []any{}
	var id int64 = -1
	if err := sqlitex.Execute(
		conn,
		query,
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				// object id
				objectID := stmt.ColumnInt64(0)
				value := d.decodeAny(stmt, 1)
				decoded, err := d.decodeFeatureValue(conn, tableFeature, value)
				if err != nil {
					return err
				}

				if id == -1 {
					id = objectID
				} else if id != objectID {
					if err := d.decodeFeatureList(id, feature, values); err != nil {
						return err
					}
					values = nil
				}
				id = objectID
				values = append(values, decoded)
				return nil
			},
		}); err != nil {
		return err
	}
	if id != -1 {
		if err := d.decodeFeatureList(id, feature, values); err != nil {
			return err
		}
	}
	return nil
}

func (d *SQLDecoder) decodeFeatureList(objectID int64, feature EStructuralFeature, values []any) error {
	if len(values) == 0 {
		return nil
	}
	eObject, _ := d.sqlIDManager.getObjectFromID(objectID)
	if eObject == nil {
		return fmt.Errorf("unable to find object with id '%v'", objectID)
	}
	eList := eObject.EGetResolve(feature, false).(EList)
	eList.AddAll(NewImmutableEList(values))
	return nil
}

func (d *SQLDecoder) addError(err error) {
	d.resource.GetErrors().Add(NewEDiagnosticImpl(err.Error(), d.resource.GetURI().String(), 0, 0))
}
