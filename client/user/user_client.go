package user

import (
	"MS_Local/pb_gen"
	"MS_Local/utils"
	"context"
	"google.golang.org/grpc"
	"log"
)

//https://www.cnblogs.com/rickiyang/p/14989375.html
type UserClient struct {
	service pb_gen.MSLocalClient
	token   utils.AuthToken
}

func NewUserClient(cc *grpc.ClientConn) *UserClient {
	cli := pb_gen.NewMSLocalClient(cc)
	return &UserClient{cli, utils.AuthToken{""}}
}

func (cli *UserClient) RegisterUser(ctx context.Context, uname string, upwd string) (*pb_gen.RegisterUserResponse, error) {
	req := &pb_gen.RegisterUserRequest{
		Name:     uname,
		Password: upwd,
	}

	res, err := cli.service.RegisterUser(ctx, req)
	if err != nil {
		log.Printf("register failed, err =  [%v]", err)
		return nil, err
	}
	log.Printf("register success, project info is \n %v", res.User)
	return res, nil
}

func (cli *UserClient) LoginUser(ctx context.Context, uname string, upwd string) (*pb_gen.LoginUserResponse, error) {
	req := &pb_gen.LoginUserRequest{
		Name:     uname,
		Password: upwd,
	}

	res, err := cli.service.LoginUser(ctx, req)
	if err != nil {
		log.Printf("login failed, err = [%v]", err)
		return nil, err
	}
	log.Printf("login success!")
	cli.token = utils.AuthToken{Token: res.Token}
	return res, err

}

func (cli *UserClient) UpdateUser(ctx context.Context, uId uint64, updateType pb_gen.UpdateType, newName string, newPwd string) (*pb_gen.UpdateUserResponse, error) {
	req := &pb_gen.UpdateUserRequest{
		Uid:         uId,
		Type:        updateType,
		NewName:     newName,
		NewPassword: newPwd,
	}
	res, err := cli.service.UpdateUser(ctx, req)
	if err != nil {
		log.Printf("update failed, err = [%v]", err)
		return nil, err
	}
	log.Printf("update %s success", res.User.Name)
	return res, nil
}

func (cli *UserClient) DeleteUser(ctx context.Context, uId uint64, name string, pwd string) (*pb_gen.DeleteUserResponse, error) {
	req := &pb_gen.DeleteUserRequest{
		Id:       uId,
		Name:     name,
		Password: pwd,
	}
	res, err := cli.service.DeleteUser(ctx, req)
	if err != nil {
		log.Printf("delete user failed, err=[%v]", err)
		return nil, err
	}
	log.Printf("delete user success")
	return res, nil
}
