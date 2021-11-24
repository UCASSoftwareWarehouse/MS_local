package routes

import (
	"MS_Local/internal/user"
	"MS_Local/pb_gen"
	"context"
)

func (s *MSLocalServer) RegisterUser(ctx context.Context, in *pb_gen.RegisterUserRequest) (*pb_gen.RegisterUserResponse, error) {
	return user.RegisterUser(ctx, in)
}
