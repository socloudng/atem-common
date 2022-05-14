package configs

type ServerConfig struct {
	Env          string `mapstructure:"env" yaml:"env"`
	ServerIp     string `mapstructure:"ip" yaml:"ip"`
	ServerPort   int    `mapstructure:"port" yaml:"port"`
	ServerName   string `mapstructure:"name" yaml:"name"`
	ReadTimeout  int    `mapstructure:"read-timeout" yaml:"read-timeout"`
	WriteTimeout int    `mapstructure:"write-timeout" yaml:"write-timeout"`
}
