package dynamicdb

import (
	"github.com/socloudng/atem-common/atem-common-middleware/db"

	"gorm.io/gorm"
)

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName string `json:"tableName" gorm:"column:table_name"`
}

type Column struct {
	DataType      string `json:"dataType" gorm:"column:data_type"`
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}

type DynamicDatabase interface {
	GetDB() (data []Db, err error)
	GetTables(dbName string) (data []Table, err error)
	GetColumn(tableName string, dbName string) (data []Column, err error)
}

var dbs map[string]DynamicDatabase = make(map[string]DynamicDatabase)

func GetDatabase(orm *gorm.DB, cfg *db.DbConfig) DynamicDatabase {
	switch cfg.DbType {
	case "mysql":
		if dbs["mysql"] == nil {
			dbs["mysql"] = NewDynamicDBMysql(orm)
		}
		return dbs["mysql"]
	case "pgsql":
		if dbs["pgsql"] == nil {
			dbs["pgsql"] = NewDynamicDBPgsql(orm, cfg)
		}
		return dbs["pgsql"]
	case "sqlserver":
		fallthrough
	case "mssql":
		if dbs["mssql"] == nil {
			dbs["mssql"] = NewDynamicDBSqlServer(orm, cfg)
		}
		return dbs["mssql"]
	case "sqlite3":
		if dbs["sqlite3"] == nil {
			dbs["sqlite3"] = NewDynamicDBSqlite3(orm, cfg)
		}
		return dbs["sqlite3"]
	default:
		if dbs["mysql"] == nil {
			dbs["mysql"] = NewDynamicDBMysql(orm)
		}
		return dbs["mysql"]
	}
}
