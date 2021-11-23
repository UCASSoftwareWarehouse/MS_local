package User

import (
	"MS_Local/mysql"
	"MS_Local/mysql/action/user"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func RegisterUser(ctx context.Context, req *pb_gen.RegisterUserRequest) (*pb_gen.RegisterUserResponse, error) {
	tmp_user := model.User{
		UserName: req.Name,
		Password: utils.Encrypt(req.Password),
	}
	_, err := user.GetUserByUserName(mysql.Mysql, tmp_user.UserName)
	if err == nil {
		log.Printf("name %s already used;", tmp_user.UserName)
		return nil, status.Errorf(codes.AlreadyExists,
			fmt.Sprintf("the name has already been used %s", tmp_user.UserName))
	}
	if len(req.Password) < 6 || len(req.Password) > 10 {
		log.Printf("the password length need to be between 6-10, len is :%d", len(req.Password))
		return nil, status.Errorf(codes.InvalidArgument,
			fmt.Sprintf("the password length need to be between 6-10"),
		)
	}
	uid, err := user.AddUser(mysql.Mysql, tmp_user)
	if err != nil {
		log.Printf("add user failed, err=[%v]", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("add failed, err=[%v]", err))
	}
	user_info := pb_gen.User{
		Id:       uid,
		Name:     req.Name,
		Password: req.Password,
	}
	return &pb_gen.RegisterUserResponse{
		User:   &user_info,
		Status: 1,
	}, nil
}
