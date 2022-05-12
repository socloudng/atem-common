package rpc

import (
	"github/socloudng/atem-common/atem-common-middleware/getcdv3"
	"log"
	"net"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type grpcServer struct {
	option            *GrpcServerConfig
	regRpcServerFuncs []func(*grpc.Server)
}

func NewGrpcServer(cfg *GrpcServerConfig) *grpcServer {
	//fixme In the configuration file, ip takes precedence, if not, get the valid network card ip of the machine
	// see https://gist.github.com/jniltinho/9787946#gistcomment-3019898

	serv := &grpcServer{option: cfg, regRpcServerFuncs: make([]func(*grpc.Server), 0)}
	serv.option.ServerIp = autoGetIp(serv.option.ServerIp)
	return serv
}

func autoGetIp(servIp string) string {
	if servIp == "" {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			panic(err.Error())
		}
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		servIp = localAddr.IP.String()
	}
	return servIp
}

func (s *grpcServer) Run(regRpcServerFunc ...func(*grpc.Server)) {
	//start network listener
	registerAddress := s.option.ServerIp + ":" + strconv.Itoa(s.option.ServerPort)
	listener, err := net.Listen("tcp", registerAddress)
	if err != nil {
		log.Panicln("0", "Listen failed ", err.Error(), registerAddress)
		return
	}
	log.Println("0", "listen network success, ", registerAddress, listener)
	defer listener.Close()

	//create grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()
	s.AddRpcFunc(regRpcServerFunc...)
	s.registRpc(srv) //regist services to grpc server
	s.registEtcd()   //regist server to etcd
	//start rpc listener
	err = srv.Serve(listener)
	if err != nil {
		log.Panicln("0", "Serve failed ", err.Error())
		return
	}

	log.Println("0", "message cms rpc  success")
}

func (s *grpcServer) AddRpcFunc(funcs ...func(*grpc.Server)) {
	if funcs != nil {
		s.regRpcServerFuncs = append(s.regRpcServerFuncs, funcs...)
	}
}

func (s *grpcServer) registRpc(srv *grpc.Server) {
	for _, fn := range s.regRpcServerFuncs {
		fn(srv) //regist server
	}
}

func (s *grpcServer) registEtcd() {
	if s.option.EtcdEnable {
		err := getcdv3.RegisterEtcd(s.option.EtcdSchema,
			strings.Join(s.option.EtcdAddrs, ","),
			s.option.ServerIp, s.option.ServerPort, s.option.ServiceName, 10)
		if err != nil {
			log.Panicln("0", "RegisterEtcd failed ", err.Error())
			return
		}
	}
}
