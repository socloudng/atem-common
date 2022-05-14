package etcd

type EtcdConfig struct {
	EtcdSchema string   `yaml:"etcdSchema"`
	EtcdAddr   []string `yaml:"etcdAddr"`
}
