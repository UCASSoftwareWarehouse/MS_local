package client

import (
	"MS_Local/config"
	"google.golang.org/grpc"
	"log"
)

func InitMSLocalClient() *grpc.ClientConn {
	conn, err := grpc.Dial(config.ServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err=[%v]", err)
		return nil
	}
	return conn
}
