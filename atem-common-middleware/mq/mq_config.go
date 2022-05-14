package mq

type MqConfig struct {

	// AsyncTransferEnable : 是否开启文件异步转移(默认同步) false
	AsyncEnable bool `mapstructure:"async-enable" yaml:"async-enable"`
	// RabbitURL : rabbitmq服务的入口url	 "amqp://guest:guest@127.0.0.1:5672/"
	MqSource string `mapstructure:"source" yaml:"source"`
}
