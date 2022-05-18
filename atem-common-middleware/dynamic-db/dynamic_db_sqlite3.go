package dynamicdb

import (
	"github.com/socloudng/atem-common/atem-common-middleware/db"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dynamicDbSqlite3 struct {
	orm    *gorm.DB
	option *db.DbConfig
}

func NewDynamicDBSqlite3(orm *gorm.DB, cfg *db.DbConfig) *dynamicDbSqlite3 {
	return &dynamicDbSqlite3{orm: orm, option: cfg}
}

func (a *dynamicDbSqlite3) GetDB() (data []Db, err error) {
	var entities []Db
	sql := `.databases`
	err = a.orm.Raw(sql).Scan(&entities).Error
	return entities, err
}

func (a *dynamicDbSqlite3) GetTables(dbName string) (data []Table, err error) {
	var entities []Table
	sql := `select name from sqlite_master where type='table' order by name`
	db, _err := a.getOrm(dbName)
	if _err != nil {
		return nil, errors.Wrapf(err, "[sqlite3] 连接 数据库(%s)的表失败!", dbName)
	}
	err = db.Raw(sql, dbName, "public").Scan(&entities).Error
	return entities, err
}

func (a *dynamicDbSqlite3) GetColumn(tableName string, dbName string) (data []Column, err error) {

	var entities []Column
	return entities, err
}

func (a *dynamicDbSqlite3) getOrm(dbName string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(a.option.Dsn(dbName)),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	return
}
