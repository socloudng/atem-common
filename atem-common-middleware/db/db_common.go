package db

import (
	"database/sql"
	"fmt"
	"log"
)

type CommonDb struct {
	db     *sql.DB
	option *DbConfig
}

func NewCommonDb(cfg *DbConfig) *CommonDb {
	cmDb := &CommonDb{option: cfg}
	db, _ := sql.Open(cfg.DbType, cfg.Dsn(""))
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to db, err:%s" + err.Error())
	}

	cmDb.db = db
	return cmDb
}

// DBConn : 返回数据库连接对象
func (c *CommonDb) DBConn() *sql.DB {
	return c.db
}

// Exec executes a query without returning any rows.
func (m CommonDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is not created")
	}
	return m.db.Exec(query, args...)
}

// Prepare creates a prepared statement for later queries or executions.
func (m *CommonDb) Prepare(query string) (*sql.Stmt, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is not created")
	}
	return m.db.Prepare(query)
}

// Query executes a query that returns rows, typically a SELECT.
func (m *CommonDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is not created")
	}
	return m.db.Query(query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
func (m *CommonDb) QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	if m.db == nil {
		return nil, fmt.Errorf("db is not created")
	}
	return m.db.QueryRow(query, args...), nil
}

// Close closes the database and prevents new queries from starting.
func (m *CommonDb) Close() error {
	if m.db == nil {
		return fmt.Errorf("db is not created")
	}
	return m.db.Close()
}

func (c *CommonDb) ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
