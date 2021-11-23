package routes

import (
	"MS_Local/internal/User"
	"MS_Local/pb_gen"
	"context"
)

func (s *MSLocalServer) AddUser(ctx context.Context, in *pb_gen.RegisterUserRequest) (*pb_gen.RegisterUserResponse, error) {
	return User.AddUser(ctx, in)
}
