package es

type ESConfig struct {
	ElasticSearchAddrs    []string `mapstructure:"address" json:"es-address" yaml:"address"`
	ElasticSearchUser     string   `mapstructure:"user" json:"es-user" yaml:"user"`
	ElasticSearchPassword string   `mapstructure:"password" json:"es-password" yaml:"password"`
}
