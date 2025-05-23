package sql

import (
	"strings"
	"sync"

	"github.com/masagroup/soft.go/ecore"
)

type SQLColumn struct {
	Table      *SQLTable
	Index      int
	ColumnName string
	ColumnType string
	Primary    bool
	Auto       bool
	Reference  *SQLTable
}

type sqlColumnOption interface {
	apply(col *SQLColumn)
}

type funcSqlColumnOption struct {
	f func(col *SQLColumn)
}

func (o *funcSqlColumnOption) apply(col *SQLColumn) {
	o.f(col)
}

func newFuncSqlColumnOption(f func(col *SQLColumn)) *funcSqlColumnOption {
	return &funcSqlColumnOption{f: f}
}

func withSqlColumnName(columnName string) sqlColumnOption {
	return newFuncSqlColumnOption(func(col *SQLColumn) {
		col.ColumnName = columnName
	})
}

func withSqlColumnPrimary(primary bool) sqlColumnOption {
	return newFuncSqlColumnOption(func(col *SQLColumn) {
		col.Primary = primary
	})
}

func withSqlColumnAuto(auto bool) sqlColumnOption {
	return newFuncSqlColumnOption(func(col *SQLColumn) {
		col.Auto = auto
	})
}

func newSqlAttributeColumn(columnName string, columnType string, options ...sqlColumnOption) *SQLColumn {
	col := &SQLColumn{
		ColumnName: columnName,
		ColumnType: columnType,
	}
	for _, opt := range options {
		opt.apply(col)
	}
	return col
}

func newSqlReferenceColumn(reference *SQLTable, options ...sqlColumnOption) *SQLColumn {
	col := &SQLColumn{
		ColumnName: reference.key.ColumnName,
		ColumnType: reference.key.ColumnType,
		Reference:  reference,
	}
	for _, opt := range options {
		opt.apply(col)
	}
	return col
}

type sqlTableOption interface {
	apply(t *SQLTable)
}

type funcSqlTableOption struct {
	f func(t *SQLTable)
}

func (o *funcSqlTableOption) apply(t *SQLTable) {
	o.f(t)
}

func newFuncSqlTableOption(f func(t *SQLTable)) *funcSqlTableOption {
	return &funcSqlTableOption{f: f}
}

func withSqlTableColumns(columns ...*SQLColumn) sqlTableOption {
	return newFuncSqlTableOption(func(t *SQLTable) {
		t.columns = columns
		for i, column := range columns {
			t.initColumn(column, i)
		}
	})
}

func withSqlTableCreateIfNotExists(createIfNotExists bool) sqlTableOption {
	return newFuncSqlTableOption(func(t *SQLTable) {
		t.createIfNotExists = createIfNotExists
	})
}

type SQLTable struct {
	name              string
	key               *SQLColumn
	columns           []*SQLColumn
	indexes           [][]*SQLColumn
	createIfNotExists bool
}

func newSqlTable(name string, options ...sqlTableOption) *SQLTable {
	t := &SQLTable{
		name: name,
	}
	for _, opt := range options {
		opt.apply(t)
	}
	return t
}

func (t *SQLTable) addColumn(column *SQLColumn) {
	t.initColumn(column, len(t.columns))
	t.columns = append(t.columns, column)
}

func (t *SQLTable) initColumn(column *SQLColumn, index int) {
	column.Table = t
	column.Index = index
	if column.Primary {
		t.key = column
	}
}

func (t *SQLTable) CreateQuery() string {
	var tableQuery strings.Builder
	tableQuery.WriteString("CREATE TABLE ")
	if t.createIfNotExists {
		tableQuery.WriteString("IF NOT EXISTS ")
	}
	tableQuery.WriteString(sqlEscapeIdentifier(t.name))
	tableQuery.WriteString(" (")
	// columns
	for i, c := range t.columns {
		if i != 0 {
			tableQuery.WriteString(",")
		}
		tableQuery.WriteString(sqlEscapeIdentifier(c.ColumnName))
		tableQuery.WriteString(" ")
		tableQuery.WriteString(c.ColumnType)
		if c.Primary {
			tableQuery.WriteString(" PRIMARY KEY")
			if c.Auto {
				tableQuery.WriteString(" AUTOINCREMENT")
			}
		}
	}
	// constraints
	for _, c := range t.columns {
		if c.Reference != nil {
			tableQuery.WriteString(",FOREIGN KEY(")
			tableQuery.WriteString(sqlEscapeIdentifier(c.ColumnName))
			tableQuery.WriteString(") REFERENCES ")
			tableQuery.WriteString(sqlEscapeIdentifier(c.Reference.name))
			tableQuery.WriteString("(")
			tableQuery.WriteString(sqlEscapeIdentifier(c.Reference.key.ColumnName))
			tableQuery.WriteString(")")
		}
	}
	tableQuery.WriteString(");")
	for _, index := range t.indexes {
		tableQuery.WriteString("\n")
		tableQuery.WriteString("CREATE INDEX ")
		if t.createIfNotExists {
			tableQuery.WriteString("IF NOT EXISTS ")
		}
		tableQuery.WriteString("\"idx_")
		tableQuery.WriteString(t.name)
		for _, c := range index {
			tableQuery.WriteString("_")
			tableQuery.WriteString(c.ColumnName)
		}
		tableQuery.WriteString("\" ON ")
		tableQuery.WriteString(sqlEscapeIdentifier(t.name))
		tableQuery.WriteString("(")
		for i, c := range index {
			if i != 0 {
				tableQuery.WriteString(",")
			}
			tableQuery.WriteString(sqlEscapeIdentifier(c.ColumnName))
		}
		tableQuery.WriteString(");")
	}
	return tableQuery.String()
}

func (t *SQLTable) InsertQuery() string {
	var tableQuery strings.Builder
	tableQuery.WriteString("INSERT INTO ")
	tableQuery.WriteString(sqlEscapeIdentifier(t.name))
	tableQuery.WriteString(" (")
	for i, c := range t.columns {
		if i != 0 {
			tableQuery.WriteString(",")
		}
		tableQuery.WriteString(sqlEscapeIdentifier(c.ColumnName))
	}
	tableQuery.WriteString(") VALUES (")
	for i := range t.columns {
		if i != 0 {
			tableQuery.WriteString(",")
		}
		tableQuery.WriteString("?")
	}
	tableQuery.WriteString(") RETURNING ")
	if t.key != nil {
		tableQuery.WriteString(sqlEscapeIdentifier(t.key.ColumnName))
	} else {
		tableQuery.WriteString("rowid")
	}
	return tableQuery.String()
}

func (t *SQLTable) InsertOrReplaceQuery() string {
	var tableQuery strings.Builder
	tableQuery.WriteString("INSERT OR REPLACE INTO ")
	tableQuery.WriteString(sqlEscapeIdentifier(t.name))
	tableQuery.WriteString(" (")
	for i, c := range t.columns {
		if i != 0 {
			tableQuery.WriteString(",")
		}
		tableQuery.WriteString(sqlEscapeIdentifier(c.ColumnName))
	}
	tableQuery.WriteString(") VALUES (")
	for i := range t.columns {
		if i != 0 {
			tableQuery.WriteString(",")
		}
		tableQuery.WriteString("?")
	}
	tableQuery.WriteString(")")
	return tableQuery.String()
}

func (t *SQLTable) DefaultValues() []any {
	return make([]any, len(t.columns))
}

func (t *SQLTable) KeyName() string {
	return sqlEscapeIdentifier(t.key.ColumnName)
}

func (t *SQLTable) SelectQuery(columns []string, selection string, orderBy string) string {
	var selectQuery strings.Builder
	selectQuery.WriteString("SELECT ")
	if len(columns) == 0 {
		selectQuery.WriteString("*")
	} else {
		for i, column := range columns {
			if i != 0 {
				selectQuery.WriteString(",")
			}
			selectQuery.WriteString(sqlEscapeIdentifier(column))
		}
	}
	selectQuery.WriteString(" FROM ")
	selectQuery.WriteString(sqlEscapeIdentifier(t.name))
	if len(selection) > 0 {
		selectQuery.WriteString(" WHERE ")
		selectQuery.WriteString(selection)
	}
	if len(orderBy) > 0 {
		selectQuery.WriteString(" ORDER BY ")
		selectQuery.WriteString(orderBy)
	}
	return selectQuery.String()
}

type SQLClassSchema struct {
	Table    *SQLTable
	Features []*SQLFeatureSchema
}

func (s *SQLClassSchema) GetFeatureSchema(feature ecore.EStructuralFeature) *SQLFeatureSchema {
	for _, f := range s.Features {
		if f.Feature == feature {
			return f
		}
	}
	return nil
}

type SQLFeatureSchema struct {
	FeatureKind SQLFeatureKind
	Feature     ecore.EStructuralFeature
	Column      *SQLColumn
	Table       *SQLTable
}

type SQLSchema struct {
	mutex             sync.Mutex
	PropertiesTable   *SQLTable
	PackagesTable     *SQLTable
	ClassesTable      *SQLTable
	ObjectsTable      *SQLTable
	ContentsTable     *SQLTable
	EnumsTable        *SQLTable
	classSchemaMap    map[ecore.EClass]*SQLClassSchema
	createIfNotExists bool
	isContainerID     bool
	objectIDName      string
}

type sqlSchemaOption interface {
	apply(s *SQLSchema)
}

type funcSqlSchemaOption struct {
	f func(s *SQLSchema)
}

func (o *funcSqlSchemaOption) apply(s *SQLSchema) {
	o.f(s)
}

func newFuncSqlSchemaOption(f func(s *SQLSchema)) *funcSqlSchemaOption {
	return &funcSqlSchemaOption{f: f}
}

func WithObjectIDName(objectIDName string) sqlSchemaOption {
	return newFuncSqlSchemaOption(func(s *SQLSchema) {
		s.objectIDName = objectIDName
	})
}

func WithContainerID(isContainerID bool) sqlSchemaOption {
	return newFuncSqlSchemaOption(func(s *SQLSchema) {
		s.isContainerID = isContainerID
	})
}

func WithCreateIfNotExists(createIfNotExists bool) sqlSchemaOption {
	return newFuncSqlSchemaOption(func(s *SQLSchema) {
		s.createIfNotExists = createIfNotExists
	})
}

func NewSqlSchema(options ...sqlSchemaOption) *SQLSchema {
	// create scheam and apply options
	s := &SQLSchema{
		classSchemaMap: map[ecore.EClass]*SQLClassSchema{},
	}
	for _, opt := range options {
		opt.apply(s)
	}

	// common tables definitions
	s.PropertiesTable = newSqlTable(
		".properties",
		withSqlTableColumns(
			newSqlAttributeColumn("key", "TEXT", withSqlColumnPrimary(true)),
			newSqlAttributeColumn("value", "TEXT"),
		),
		withSqlTableCreateIfNotExists(s.createIfNotExists),
	)
	s.PackagesTable = newSqlTable(
		".packages",
		withSqlTableColumns(
			newSqlAttributeColumn("packageID", "INTEGER", withSqlColumnPrimary(true), withSqlColumnAuto(true)),
			newSqlAttributeColumn("uri", "TEXT"),
		),
		withSqlTableCreateIfNotExists(s.createIfNotExists),
	)
	s.ClassesTable = newSqlTable(
		".classes",
		withSqlTableColumns(
			newSqlAttributeColumn("classID", "INTEGER", withSqlColumnPrimary(true), withSqlColumnAuto(true)),
			newSqlReferenceColumn(s.PackagesTable),
			newSqlAttributeColumn("name", "TEXT"),
		),
		withSqlTableCreateIfNotExists(s.createIfNotExists),
	)
	s.ObjectsTable = newSqlTable(
		".objects",
		withSqlTableColumns(
			newSqlAttributeColumn("objectID", "INTEGER", withSqlColumnPrimary(true), withSqlColumnAuto(true)),
			newSqlReferenceColumn(s.ClassesTable),
		),
		withSqlTableCreateIfNotExists(s.createIfNotExists),
	)
	// container and feayure id in objects table
	if s.isContainerID {
		s.ObjectsTable.addColumn(newSqlReferenceColumn(s.ObjectsTable, withSqlColumnName("containerID")))
		s.ObjectsTable.addColumn(newSqlAttributeColumn("containerFeatureID", "INTEGER"))
	}
	// add id attribute column if name is not object table primary key
	if len(s.objectIDName) > 0 && s.objectIDName != s.ObjectsTable.key.ColumnName {
		s.ObjectsTable.addColumn(newSqlAttributeColumn(s.objectIDName, "TEXT"))
	}
	s.ContentsTable = newSqlTable(
		".contents",
		withSqlTableColumns(
			newSqlReferenceColumn(s.ObjectsTable),
		),
		withSqlTableCreateIfNotExists(s.createIfNotExists),
	)
	s.EnumsTable = newSqlTable(
		".enums",
		withSqlTableColumns(
			newSqlAttributeColumn("enumID", "INTEGER", withSqlColumnPrimary(true), withSqlColumnAuto(true)),
			newSqlReferenceColumn(s.PackagesTable),
			newSqlAttributeColumn("name", "TEXT"),
			newSqlAttributeColumn("literal", "TEXT"),
		),
		withSqlTableCreateIfNotExists(s.createIfNotExists),
	)
	return s
}

func (s *SQLSchema) GetClassSchema(eClass ecore.EClass) *SQLClassSchema {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.getOrComputeClassSchema(eClass)
}

func (s *SQLSchema) getOrComputeClassSchema(eClass ecore.EClass) *SQLClassSchema {
	classSchema := s.classSchemaMap[eClass]
	if classSchema == nil {
		// create table descriptor
		classTable := newSqlTable(strings.ToLower(eClass.GetName()), withSqlTableCreateIfNotExists(s.createIfNotExists))
		classTable.addColumn(newSqlAttributeColumn(strings.ToLower(eClass.GetName())+"ID", "INTEGER", withSqlColumnPrimary(true)))

		// compute table columns and external tables
		eFeatures := eClass.GetEStructuralFeatures()
		classSchema = &SQLClassSchema{
			Table:    classTable,
			Features: make([]*SQLFeatureSchema, 0, eFeatures.Size()),
		}

		// register class data now to handle correctly cycles references
		s.classSchemaMap[eClass] = classSchema

		newFeatureReferenceColumn := func(featureSchema *SQLFeatureSchema, eFeature ecore.EStructuralFeature, table *SQLTable) {
			column := newSqlReferenceColumn(table, withSqlColumnName(eFeature.GetName()))
			classTable.addColumn(column)
			featureSchema.Column = column
		}

		newFeatureAttributeColumn := func(featureSchema *SQLFeatureSchema, eFeature ecore.EStructuralFeature, columnType string) {
			column := newSqlAttributeColumn(eFeature.GetName(), columnType)
			classTable.addColumn(column)
			featureSchema.Column = column
		}

		newFeatureTable := func(featureSchema *SQLFeatureSchema, eFeature ecore.EStructuralFeature, columns ...*SQLColumn) {
			table := newSqlTable(
				classTable.name+"_"+eFeature.GetName(),
				withSqlTableColumns(columns...),
				withSqlTableCreateIfNotExists(s.createIfNotExists),
			)
			table.key = columns[0]
			table.indexes = [][]*SQLColumn{{columns[0], columns[1]}}
			featureSchema.Table = table
		}

		for itFeature := eFeatures.Iterator(); itFeature.HasNext(); {
			eFeature := itFeature.Next().(ecore.EStructuralFeature)
			// new feature data
			featureSchema := &SQLFeatureSchema{
				Feature:     eFeature,
				FeatureKind: GetSQLCodecFeatureKind(eFeature),
			}
			classSchema.Features = append(classSchema.Features, featureSchema)

			// compute class table columns or children tables
			switch featureSchema.FeatureKind {
			case sfkObject:
				// retrieve object reference type
				eReference := eFeature.(ecore.EReference)
				referenceSchema := s.getOrComputeClassSchema(eReference.GetEReferenceType())
				newFeatureReferenceColumn(featureSchema, eFeature, referenceSchema.Table)
			case sfkObjectReference:
				newFeatureAttributeColumn(featureSchema, eFeature, "TEXT")
			case sfkObjectList:
				// internal reference
				eReference := eFeature.(ecore.EReference)
				referenceSchema := s.getOrComputeClassSchema(eReference.GetEReferenceType())
				newFeatureTable(featureSchema, eFeature,
					newSqlReferenceColumn(classTable),
					newSqlAttributeColumn("idx", "REAL"),
					newSqlReferenceColumn(referenceSchema.Table, withSqlColumnName(eFeature.GetName())),
				)
			case sfkObjectReferenceList:
				newFeatureTable(featureSchema, eFeature,
					newSqlReferenceColumn(classTable),
					newSqlAttributeColumn("idx", "REAL"),
					newSqlAttributeColumn("uri", "TEXT"),
				)
			case sfkEnum:
				newFeatureReferenceColumn(featureSchema, eFeature, s.EnumsTable)
			case sfkFloat64, sfkFloat32:
				newFeatureAttributeColumn(featureSchema, eFeature, "REAL")
			case sfkBool, sfkByte, sfkInt, sfkInt16, sfkInt32, sfkInt64:
				newFeatureAttributeColumn(featureSchema, eFeature, "INTEGER")
			case sfkDate, sfkString, sfkData:
				newFeatureAttributeColumn(featureSchema, eFeature, "TEXT")
			case sfkByteArray:
				newFeatureAttributeColumn(featureSchema, eFeature, "BLOB")
			case sfkDataList:
				newFeatureTable(featureSchema, eFeature,
					newSqlReferenceColumn(classTable),
					newSqlAttributeColumn("idx", "REAL"),
					newSqlAttributeColumn(eFeature.GetName(), "TEXT"),
				)
			}
		}
	}
	return classSchema
}

var sqliteKeywWords = []string{
	"ABORT", "ACTION", "ADD", "AFTER", "ALL", "ALTER", "ALWAYS", "ANALYZE", "AND", "AS", "ASC", "ATTACH",
	"AUTOINCREMENT", "BEFORE", "BEGIN", "BETWEEN", "BY", "CASCADE", "CASE", "CAST", "CHECK", "COLLATE",
	"COLUMN", "COMMIT", "CONFLICT", "CONSTRAINT", "CREATE", "CROSS", "CURRENT", "CURRENT_DATE",
	"CURRENT_TIME", "CURRENT_TIMESTAMP", "DATABASE", "DEFAULT", "DEFERRABLE", "DEFERRED", "DELETE",
	"DESC", "DETACH", "DISTINCT", "DO", "DROP", "EACH", "ELSE", "END", "ESCAPE", "EXCEPT", "EXCLUDE",
	"EXCLUSIVE", "EXISTS", "EXPLAIN", "FAIL", "FILTER", "FIRST", "FOLLOWING", "FOR", "FOREIGN", "FROM",
	"FULL", "GENERATED", "GLOB", "GROUP", "GROUPS", "HAVING", "IF", "IGNORE", "IMMEDIATE", "IN", "INDEX",
	"INDEXED", "INITIALLY", "INNER", "INSERT", "INSTEAD", "INTERSECT", "INTO", "IS", "ISNULL", "JOIN",
	"KEY", "LAST", "LEFT", "LIKE", "LIMIT", "MATCH", "MATERIALIZED", "NATURAL", "NO", "NOT", "NOTHING",
	"NOTNULL", "NULL", "NULLS", "OF", "OFFSET", "ON", "OR", "ORDER", "OTHERS", "OUTER", "OVER", "PARTITION",
	"PLAN", "PRAGMA", "PRECEDING", "PRIMARY", "QUERY", "RAISE", "RANGE", "RECURSIVE", "REFERENCES", "REGEXP",
	"REINDEX", "RELEASE", "RENAME", "REPLACE", "RESTRICT", "RETURNING", "RIGHT", "ROLLBACK", "ROW", "ROWS", "SAVEPOINT",
	"SELECT", "SET", "TABLE", "TEMP", "TEMPORARY", "THEN", "TIES", "TO", "TRANSACTION", "TRIGGER", "UNBOUNDED", "UNION",
	"UNIQUE", "UPDATE", "USING", "VACUUM", "VALUES", "VIEW", "VIRTUAL", "WHEN", "WHERE", "WINDOW", "WITH", "WITHOUT",
}

var sqliteKeyWordsAsSet = func() (result map[string]struct{}) {
	result = map[string]struct{}{}
	for _, keyword := range sqliteKeywWords {
		result[keyword] = struct{}{}
	}
	return
}()

func isSqliteKeyWord(k string) bool {
	if _, isKeyWord := sqliteKeyWordsAsSet[k]; isKeyWord {
		return true
	}
	if _, isKeyWord := sqliteKeyWordsAsSet[strings.ToUpper(k)]; isKeyWord {
		return true
	}
	return false
}

func sqlEscapeIdentifier(id string) string {
	if id[0] == '.' || isSqliteKeyWord(id) {
		return "\"" + id + "\""
	}
	return id
}
