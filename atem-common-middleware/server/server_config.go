package server

type ServerConfig struct {
	ServerEnv     string `mapstructure:"env" yaml:"env"`
	ServerIp      string `mapstructure:"ip" yaml:"ip"`
	ServerPort    int    `mapstructure:"port" yaml:"port"`
	ServerName    string `mapstructure:"name" yaml:"name"`
	ReadTimeout   int    `mapstructure:"read-timeout" yaml:"read-timeout"`
	WriteTimeout  int    `mapstructure:"write-timeout" yaml:"write-timeout"`
	HttpsEnabled  bool   `mapstructure:"https" yaml:"https"`
	HttpsCertPath string `mapstructure:"cert-path" yaml:"cert-path"`
	HttpsKeyPath  string `mapstructure:"key-path" yaml:"key-path"`
}
