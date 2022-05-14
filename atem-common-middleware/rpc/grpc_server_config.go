package rpc

import (
	"github.com/socloudng/atem-common/atem-common-middleware/getcdv3"
	"github.com/socloudng/atem-common/configs"
)

type GrpcServerConfig struct {
	configs.ServerConfig `mapstructure:"server" yaml:"server"`
	getcdv3.EtcdConfig   `mapstructure:"etcd" yaml:"etcd"`
}
