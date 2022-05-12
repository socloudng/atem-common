package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	db     *sql.DB
	option *DbConfig
}

func NewMysqlDb(config *DbConfig) *MysqlDb {
	mysqlDb := &MysqlDb{option: config}

	db, _ := sql.Open("mysql", config.Dsn(""))
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to mysql, err:%s" + err.Error())
	}

	mysqlDb.db = db
	return mysqlDb
}

// DBConn : 返回数据库连接对象
func (c *MysqlDb) DBConn() *sql.DB {
	return c.db
}

func (c *MysqlDb) ParseRows(rows *sql.Rows) []map[string]interface{} {
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
