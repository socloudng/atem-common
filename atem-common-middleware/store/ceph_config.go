package store

type CephConfig struct {
	// CephAccessKey : 访问Key		"PFEA7NXWXSOWVTFA16C9"
	CephAccessKey string `mapstructure:"access-key" yaml:"access-key"`
	// CephSecretKey : 访问密钥		"cf3dwPMeadGbtEgwFUEA6emRVrVfDHpv0pLXFYby"
	CephSecretKey string `mapstructure:"secret-key" yaml:"secret-key"`
	// CephGWEndpoint : gateway地址	"http://<你的rgw_host>:<<你的rgw_port>>"
	CephGWEndpoint string `mapstructure:"endpoint" yaml:"endpoint"`
}
