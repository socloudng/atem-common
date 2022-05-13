package rpc

import (
	"github/socloudng/atem-common/atem-common-middleware/getcdv3"
	"github/socloudng/atem-common/configs"
)

type GrpcServerConfig struct {
	configs.ServerConfig `mapstructure:"server" yaml:"server"`
	getcdv3.EtcdConfig   `mapstructure:"etcd" yaml:"etcd"`
}
