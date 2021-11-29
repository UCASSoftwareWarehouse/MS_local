package main

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"MS_Local/pb_gen"
	"MS_Local/server/routes"
	"google.golang.org/grpc"

	"log"
	"net"
)

const (
	port = ":50051"
)

//type server struct {
//	pb_gen.UnimplementedMSLocalServer
//}

//func (s *server) SayHello(ctx context.Context, in *pb_gen.HelloRequest) (*pb_gen.HelloReply, error) {
//	log.Printf("Received: %v", in.GetName())
//	return &pb_gen.HelloReply{
//			Message: "Hello " + in.GetName(),
//		},
//		nil
//}

func main() {
	config.InitConfig()
	err := mysql.InitMysql()
	if err != nil {
		log.Fatalf("init mysql error:[%v]", err)
	}
	err = mongodb.InitMongo()
	if err != nil {
		log.Fatalf("init mongodb error:[%v]", err)
	}
	log.Printf("init mysql success!")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb_gen.RegisterMSLocalServer(s, routes.NewMSLocalServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
