package cache

type RedisConfig struct {
	DB   int    `mapstructure:"db" yaml:"db"`     // redis的哪个数据库
	Addr string `mapstructure:"addr" yaml:"addr"` // 服务器地址:端口
	Pass string `mapstructure:"auth" yaml:"auth"` // 密码
}
