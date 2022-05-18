package db

import "fmt"

type DbConfig struct {
	DbType       string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                     // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // 端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	DbName       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
}

func (m *DbConfig) Dsn(dbname string) string {
	if dbname == "" {
		dbname = m.DbName
	}
	switch dbname {
	case "mysql":
		return m.mysqlDsn(dbname)
	case "pgsql":
		return m.pgsqlDsn(dbname)
	case "sqlite3":
		return m.sqlite3Dsn(dbname)
	case "sqlserver":
		fallthrough
	case "mssql":
		return m.mssqlDsn(dbname)
	case "clickhouse":
		return m.clickhouseDsn(dbname)
	default:
		return m.mysqlDsn(dbname)
	}
}

func (m *DbConfig) mysqlDsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		m.Username, m.Password, m.Path, m.Port, dbname) + "?" + m.Config
}

func (p *DbConfig) pgsqlDsn(dbname string) string {
	return fmt.Sprintf("user=%s password=%s port=%s host=%s dbname=%s ",
		p.Username, p.Password, p.Path, p.Port, dbname) + p.Config
}

func (m *DbConfig) mssqlDsn(dbname string) string {
	// github.com/denisenkom/go-mssqldb
	return fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", m.Username, m.Password, m.Path, m.Port, dbname)
}

func (c *DbConfig) clickhouseDsn(dbname string) string {
	dsn := fmt.Sprintf("tcp://%v:%v?username=%v&password=%v&database=%v",
		c.Username, c.Password, c.Path, c.Port, dbname)
	if c.Config != "" {
		dsn = dsn + "&" + c.Config
	}
	return dsn
}

func (s *DbConfig) sqlite3Dsn(dbname string) string {

	// github.com/mattn/go-sqlite3
	return s.Path
}
