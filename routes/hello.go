package routes

import (
	"MS_Local/internal/hello"
	"MS_Local/pb_gen"
	"context"
)

type MSLocalServer struct {
	*pb_gen.UnimplementedMSLocalServer
}

func NewMSLocalServer() *MSLocalServer {
	return &MSLocalServer{}
}

func (s *MSLocalServer) SayHello(ctx context.Context, in *pb_gen.HelloRequest) (*pb_gen.HelloReply, error) {
	return hello.SayHello(ctx, in)
}
