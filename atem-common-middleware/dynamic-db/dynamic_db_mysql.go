package dynamicdb

import (
	"gorm.io/gorm"
)

type dynamicDbMysql struct {
	orm *gorm.DB
}

func NewDynamicDBMysql(orm *gorm.DB) *dynamicDbMysql {
	return &dynamicDbMysql{orm: orm}
}

func (s *dynamicDbMysql) GetDB() (data []Db, err error) {
	var entities []Db
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;"
	err = s.orm.Raw(sql).Scan(&entities).Error
	return entities, err
}

func (s *dynamicDbMysql) GetTables(dbName string) (data []Table, err error) {
	var entities []Table
	sql := `select table_name as table_name from information_schema.tables where table_schema = ?`
	err = s.orm.Raw(sql, dbName).Scan(&entities).Error
	return entities, err
}

func (s *dynamicDbMysql) GetColumn(tableName string, dbName string) (data []Column, err error) {
	var entities []Column
	sql := `
	SELECT COLUMN_NAME        column_name,
       DATA_TYPE          data_type,
       CASE DATA_TYPE
           WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH
           WHEN 'double' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'decimal' THEN CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
           WHEN 'int' THEN c.NUMERIC_PRECISION
           WHEN 'bigint' THEN c.NUMERIC_PRECISION
           ELSE '' END AS data_type_long,
       COLUMN_COMMENT     column_comment
	FROM INFORMATION_SCHEMA.COLUMNS c
	WHERE table_name = ?
	  AND table_schema = ?
	`
	err = s.orm.Raw(sql, tableName, dbName).Scan(&entities).Error
	return entities, err
}
