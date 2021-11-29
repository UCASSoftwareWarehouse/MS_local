package user

import (
	project2 "MS_Local/internal/project"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
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
		User: &user_info,
	}, nil
}

func LoginUser(ctx context.Context, req *pb_gen.LoginUserRequest) (*pb_gen.LoginUserResponse, error) {
	//uname := req.Name
	uinfo, err := user.GetUserByUserName(mysql.Mysql, req.Name)
	if err != nil {
		return nil, err
	}
	if utils.Encrypt(req.Password) != uinfo.Password {
		log.Printf("wrong password")
		return nil, status.Errorf(codes.InvalidArgument, "wrong password")
	}
	token, err := utils.CreateToken(uinfo.UserName)
	if err != nil {
		return nil, err
	}
	log.Printf("user %s login in", req.Name)
	return &pb_gen.LoginUserResponse{Token: token}, nil
}

func DeleteUser(ctx context.Context, req *pb_gen.DeleteUserRequest) (*pb_gen.DeleteUserResponse, error) {
	auth, err := utils.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("delete user %s", auth)
	uinfo, err := user.GetUserByUserId(mysql.Mysql, req.Id)
	if err != nil {
		return nil, err
	}
	if uinfo.UserName != req.Name {
		return nil, status.Errorf(codes.InvalidArgument, "the name is wrong")
	}
	if uinfo.Password != utils.Encrypt(req.Password) {
		return nil, status.Errorf(codes.InvalidArgument, "the password is wrong")
	}
	//delete projects
	var projects []model.Project
	//删除所有project
	err = project.GetProjectsByUserId(mysql.Mysql, req.Id, 10, 1, &projects)
	for _, p := range projects {
		err = project2.DeleteProject(p.ID)
		if err != nil {
			return nil, err
		}
	}
	// delete user
	err = user.DeleteByUserId(mysql.Mysql, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb_gen.DeleteUserResponse{Message: "success"}, nil
}

func UpdateUser(ctx context.Context, req *pb_gen.UpdateUserRequest) (*pb_gen.UpdateUserResponse, error) {
	auth, err := utils.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("update user %s", auth)
	uinfo, err := user.GetUserByUserId(mysql.Mysql, req.Uid)
	if err != nil {
		return nil, err
	}
	//if uinfo.UserName != req.User.Name {
	//	return nil, status.Errorf(codes.InvalidArgument, "the name is wrong")
	//}
	//if uinfo.Password != utils.Encrypt(req.User.Password) {
	//	return nil, status.Errorf(codes.InvalidArgument, "the password is wrong")
	//}

	if req.Type == pb_gen.UpdateType_uname || req.Type == pb_gen.UpdateType_all {
		err = user.UpdateUserInfo(mysql.Mysql, uinfo.ID, model.UserColumns.UserName, req.NewName)
		if err != nil {
			return nil, err
		}
	}
	if req.Type == pb_gen.UpdateType_upwd || req.Type == pb_gen.UpdateType_all {
		err = user.UpdateUserInfo(mysql.Mysql, uinfo.ID, model.UserColumns.Password, utils.Encrypt(req.NewPassword))
		if err != nil {
			return nil, err
		}
	}

	return &pb_gen.UpdateUserResponse{User: &pb_gen.User{
		Id:       uinfo.ID,
		Name:     uinfo.UserName,
		Password: uinfo.Password,
	},
	}, nil
}
