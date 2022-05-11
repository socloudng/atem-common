package rpc

import (
	"atem/atem-common/atem-common-middleware/getcdv3"
	"atem/atem-common/configs"
)

type GrpcServerConfig struct {
	configs.ServerConfig `mapstructure:"server" yaml:"server"`
	getcdv3.EtcdConfig   `mapstructure:"etcd" yaml:"etcd"`
}
