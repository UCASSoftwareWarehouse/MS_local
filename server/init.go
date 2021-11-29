package server

import (
	"MS_Local/config"
	"MS_Local/pb_gen"
	"MS_Local/server/routes"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartServe() {
	appName := config.Conf.AppName
	addr := fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port)
	log.Printf("%s Dialing addr: %s", appName, addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var options []grpc.ServerOption
	options = append(options, grpc.MaxSendMsgSize(5*1024*1024*1024*1024), grpc.MaxRecvMsgSize(5*1024*1024*1024*1024))
	grpcServer := grpc.NewServer(options...)
	//consul.MustRegisterGRPCServer(grpcServer)
	pb_gen.RegisterMSLocalServer(grpcServer, routes.NewMSLocalServer())
	log.Printf("%s ready to server at %s...", appName, addr)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer Serve failed, err=[%v]", err)
	}
}
