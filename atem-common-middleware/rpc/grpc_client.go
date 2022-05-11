package rpc

import (
	"atem/atem-common/atem-common-middleware/getcdv3"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type grpcClient struct {
	option *GrpcServerConfig
}

func NewGrpcClient(cfg *GrpcServerConfig) *grpcClient {
	return &grpcClient{option: cfg}
}

func (cli *grpcClient) getGrpcConnOfEtcdPoolService(serviceName string) *grpc.ClientConn {
	etcdPool, err := getcdv3.GetConnPool(cli.option.EtcdSchema, strings.Join(cli.option.EtcdAddrs, ","), serviceName)
	if err != nil {
		log.Fatalln("连接 gPRC 服务失败,", err)
	}
	return etcdPool.ClientConn
}

func (cli *grpcClient) getGrpcConnOfEtcdService(serviceName string) *grpc.ClientConn {
	etcdConn := getcdv3.GetConn(cli.option.EtcdSchema, strings.Join(cli.option.EtcdAddrs, ","), serviceName)
	return etcdConn
}

func (cli *grpcClient) getGrpcConnDirect() *grpc.ClientConn {
	conn, err := grpc.Dial(cli.option.ServerIp+":"+strconv.Itoa(cli.option.ServerPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalln("连接 gPRC 服务失败,", err)
	}
	return conn
}

func (cli *grpcClient) GetGrpcConn() *grpc.ClientConn {
	if cli.option.EtcdEnable {
		if cli.option.EtcdUsePool {
			return cli.getGrpcConnOfEtcdPoolService(cli.option.ServiceName)
		}
		return cli.getGrpcConnOfEtcdService(cli.option.ServiceName)
	} else {
		return cli.getGrpcConnDirect()
	}
}
