package routes

import (
	"MS_Local/pb_gen"
	"context"
	"log"
)

type MSLocalServer struct {
	*pb_gen.UnimplementedMSLocalServer
}

func NewMSLocalServer() *MSLocalServer {
	return &MSLocalServer{}
}

func (s *MSLocalServer) SayHello(ctx context.Context, in *pb_gen.HelloRequest) (*pb_gen.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb_gen.HelloReply{
			Message: "Hello " + in.GetName(),
		},
		nil
}
