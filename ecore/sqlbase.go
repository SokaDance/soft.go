package ecore

import (
	"context"
	"database/sql"
	"fmt"
	"maps"
	"reflect"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/require"
)

type sqlSafeDB struct {
	delegate *sql.DB
	sync.Mutex
}

func (db *sqlSafeDB) Conn(ctx context.Context) (*sql.Conn, error) {
	db.Lock()
	defer db.Unlock()
	return db.delegate.Conn(ctx)
}

func (db *sqlSafeDB) Prepare(query string) (*sqlSafeStmt, error) {
	db.Lock()
	defer db.Unlock()
	stmt, err := db.delegate.PrepareContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	return &sqlSafeStmt{delegate: stmt, db: db}, nil
}

func (db *sqlSafeDB) Exec(query string, args ...any) (sql.Result, error) {
	db.Lock()
	defer db.Unlock()
	return db.delegate.ExecContext(context.Background(), query, args...)
}

func (db *sqlSafeDB) Query(query string, args ...any) (*sql.Rows, error) {
	db.Lock()
	defer db.Unlock()
	return db.delegate.QueryContext(context.Background(), query, args...)
}

func (db *sqlSafeDB) QueryRow(query string, args ...any) *sql.Row {
	db.Lock()
	defer db.Unlock()
	return db.delegate.QueryRowContext(context.Background(), query, args...)
}

func (db *sqlSafeDB) Begin() (*sqlSafeTx, error) {
	db.Lock()
	defer db.Unlock()
	tx, err := db.delegate.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &sqlSafeTx{delegate: tx, db: db}, nil
}

func newSQLSafeDB(db *sql.DB) *sqlSafeDB {
	return &sqlSafeDB{delegate: db}
}

type sqlSafeStmt struct {
	delegate *sql.Stmt
	db       *sqlSafeDB
}

func (stmt *sqlSafeStmt) Exec(args ...any) (sql.Result, error) {
	stmt.db.Lock()
	defer stmt.db.Unlock()
	return stmt.delegate.ExecContext(context.Background(), args...)
}

func (stmt *sqlSafeStmt) Query(args ...any) (*sql.Rows, error) {
	stmt.db.Lock()
	defer stmt.db.Unlock()
	return stmt.delegate.QueryContext(context.Background(), args...)
}

func (stmt *sqlSafeStmt) QueryRow(args ...any) *sql.Row {
	stmt.db.Lock()
	defer stmt.db.Unlock()
	return stmt.delegate.QueryRowContext(context.Background(), args...)
}

type sqlSafeTx struct {
	delegate *sql.Tx
	db       *sqlSafeDB
}

func (tx *sqlSafeTx) Commit() error {
	tx.db.Lock()
	defer tx.db.Unlock()
	return tx.delegate.Commit()
}

func (tx *sqlSafeTx) Rollback() error {
	tx.db.Lock()
	defer tx.db.Unlock()
	return tx.delegate.Rollback()
}

func (tx *sqlSafeTx) Stmt(stmt *sqlSafeStmt) *sqlSafeStmt {
	tx.db.Lock()
	defer tx.db.Unlock()
	return &sqlSafeStmt{delegate: tx.delegate.Stmt(stmt.delegate), db: tx.db}
}

type sqlBase struct {
	db              *sqlSafeDB
	schema          *sqlSchema
	uri             *URI
	idAttributeName string
	idManager       EObjectIDManager
}

func (s *sqlBase) encodeProperties() error {
	// properties
	propertiesQuery := `
	PRAGMA synchronous = NORMAL;
	PRAGMA journal_mode = WAL;
	`
	_, err := s.db.Exec(propertiesQuery)
	return err
}

func (s *sqlBase) encodeSchema() error {
	// tables
	for _, table := range []*sqlTable{
		s.schema.packagesTable,
		s.schema.classesTable,
		s.schema.objectsTable,
		s.schema.contentsTable,
		s.schema.enumsTable,
	} {
		if _, err := s.db.Exec(table.createQuery()); err != nil {
			return err
		}
	}
	return nil
}

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	DisableMethods:          true,
	MaxDepth:                10,
}

func getDBTables(db *sql.DB) (map[string]struct{}, error) {
	rows, err := db.Query("SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := map[string]struct{}{}
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables[table] = struct{}{}
	}
	return tables, nil
}

func getDBRows(db *sql.DB, table string) ([][]any, error) {
	rows, err := db.Query("SELECT rowid,* FROM " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// get column type info
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	// used for allocation & dereferencing
	rowValues := make([]reflect.Value, len(columnTypes))
	for i := 0; i < len(columnTypes); i++ {
		// allocate reflect.Value representing a **T value
		rowValues[i] = reflect.New(reflect.PtrTo(columnTypes[i].ScanType()))
	}

	resultList := [][]any{}
	for rows.Next() {
		// initially will hold pointers for Scan, after scanning the
		// pointers will be dereferenced so that the slice holds actual values
		rowResult := make([]any, len(columnTypes))
		for i := 0; i < len(columnTypes); i++ {
			// get the **T value from the reflect.Value
			rowResult[i] = rowValues[i].Interface()
		}

		// scan each column value into the corresponding **T value
		if err := rows.Scan(rowResult...); err != nil {
			return nil, err
		}

		// dereference pointers
		for i := 0; i < len(rowValues); i++ {
			// first pointer deref to get reflect.Value representing a *T value,
			// if rv.IsNil it means column value was NULL
			if rv := rowValues[i].Elem(); rv.IsNil() {
				rowResult[i] = nil
			} else {
				// second deref to get reflect.Value representing the T value
				// and call Interface to get T value from the reflect.Value
				rowResult[i] = rv.Elem().Interface()
			}
		}

		resultList = append(resultList, rowResult)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return resultList, nil
}

func RequireEqualDB(t require.TestingT, expected, actual *sql.DB) {
	te, err := getDBTables(expected)
	if err != nil {
		require.Fail(t, err.Error())
	}
	ta, err := getDBTables(actual)
	if err != nil {
		require.Fail(t, err.Error())
	}
	if !maps.Equal(te, ta) {
		e := spewConfig.Sdump(te)
		a := spewConfig.Sdump(ta)
		diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
			A:        difflib.SplitLines(e),
			B:        difflib.SplitLines(a),
			FromFile: "Expected",
			FromDate: "",
			ToFile:   "Actual",
			ToDate:   "",
			Context:  1,
		})
		require.Fail(t, diff)
	}
	for table := range te {
		e, err := getDBRows(expected, table)
		if err != nil {
			require.Fail(t, err.Error())
		}
		a, err := getDBRows(actual, table)
		if err != nil {
			require.Fail(t, err.Error())
		}
		require.Equal(t, e, a, fmt.Sprintf("rows for tables '%s' are different", table))
	}
}
