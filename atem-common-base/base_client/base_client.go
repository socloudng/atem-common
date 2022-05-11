package base_client

import (
	"atem/atem-common/atem-common-middleware/rpc"

	"google.golang.org/grpc"
)

type baseClient struct {
	option *rpc.GrpcServerConfig
}

func (c *baseClient) conn() *grpc.ClientConn {
	conn := rpc.NewGrpcClient(c.option).GetGrpcConn()
	return conn
}

func (c *baseClient) SetConfig(config *rpc.GrpcServerConfig) *baseClient {
	c.option = config
	return c
}
