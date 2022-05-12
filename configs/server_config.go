package configs

type ServerConfig struct {
	ServerIp   string `mapstructure:"ip" yaml:"ip"`
	ServerPort int    `mapstructure:"port" yaml:"port"`
}
