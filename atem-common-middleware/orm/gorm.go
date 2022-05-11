package orm

import (
	"atem/atem-common/atem-common-middleware/db"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm(cfg *db.DbConfig) *gorm.DB {
	switch cfg.DbType {
	case "mysql":
		return GormMysql(cfg)
	case "pgsql":
		return GormPgSql(cfg)
	case "sqlite3":
		return GormMysql(cfg)
	default:
		return GormMysql(cfg)
	}
}

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func GormMysql(m *db.DbConfig) *gorm.DB {
	if m.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(m.DbName), // DSN data source name
		DefaultStringSize:         191,             // string 类型字段的默认长度
		SkipInitializeWithVersion: false,           // 根据版本自动配置
	}
	if orm, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := orm.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return orm
	}
}
func GormPgSql(p *db.DbConfig) *gorm.DB {
	if p.DbName == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(p.DbName), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if orm, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := orm.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return orm
	}
}
