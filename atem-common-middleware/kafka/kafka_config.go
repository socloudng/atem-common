package kafka

type KafkaConfig struct {
	Ws2mschat struct {
		Addr  []string `yaml:"addr"`
		Topic string   `yaml:"topic"`
	}
	Ms2pschat struct {
		Addr  []string `yaml:"addr"`
		Topic string   `yaml:"topic"`
	}
	ConsumerGroupID struct {
		MsgToMongo string `yaml:"msgToMongo"`
		MsgToMySql string `yaml:"msgToMySql"`
		MsgToPush  string `yaml:"msgToPush"`
	}
}
