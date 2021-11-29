package routes

import (
	"MS_Local/internal/user"
	"MS_Local/pb_gen"
	"context"
)

func (s *MSLocalServer) RegisterUser(ctx context.Context, in *pb_gen.RegisterUserRequest) (*pb_gen.RegisterUserResponse, error) {
	return user.RegisterUser(ctx, in)
}

func (s *MSLocalServer) DeleteUser(ctx context.Context, in *pb_gen.DeleteUserRequest) (*pb_gen.DeleteUserResponse, error) {
	return user.DeleteUser(ctx, in)
}

func (s *MSLocalServer) UpdateUser(ctx context.Context, in *pb_gen.UpdateUserRequest) (*pb_gen.UpdateUserResponse, error) {
	return user.UpdateUser(ctx, in)
}

func (s *MSLocalServer) LoginUser(ctx context.Context, in *pb_gen.LoginUserRequest) (*pb_gen.LoginUserResponse, error) {
	return user.LoginUser(ctx, in)
}
