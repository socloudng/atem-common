package store

/*
 Minio
	endpoint : 172.31.102.222:9000
	accessKeyID : admin
	secretAccessKey : test123
	useSSL : false

 Aliyun
	bucket: buckettest-filestore2
	endpoint: oss-cn-shenzhen.aliyuncs.com
	access-key: access-key
	access-secret: <你的AccessKeySecret>
	use-ssl: false
*/
type StoreConfig struct {
	StoreType       string `mapstructure:"type" yaml:"type"`
	LocalPath       string `mapstructure:"path" yaml:"path"`
	OSSBucket       string `mapstructure:"bucket" yaml:"bucket"`
	OSSEndpoint     string `mapstructure:"endpoint" yaml:"endpoint"`
	OSSAccesskey    string `mapstructure:"access-key" yaml:"access-key"`
	OSSAccessSecret string `mapstructure:"access-secret" yaml:"access-secret"`
	UseSLL          bool   `mapstructure:"use-ssl" yaml:"use-ssl"`
}
