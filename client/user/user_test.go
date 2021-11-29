package user

import (
	"MS_Local/config"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestUserClient_UpdateUser(t *testing.T) {
	conn, err := grpc.Dial(config.ServerAddr, grpc.WithInsecure())
	if err != nil {
		t.Errorf("err=[%v]", err)
	}
	defer conn.Close()
	cli := NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//// register
	//res1, err:= cli.RegisterUser(ctx,"guohao", "123456")
	//if err!=nil{
	//	t.Errorf("register user failed, err=[%v]", err)
	//}
	//login
	res2, err := cli.LoginUser(ctx, "guohao", "123456")
	if err != nil {
		t.Errorf("login failed, err=[%v]", err)
	}
	requestToken := new(utils.AuthToken)
	requestToken.Token = res2.Token
	//update user
	conn, err = grpc.Dial(config.ServerAddr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		log.Printf("failed to connect: %v", err)
		t.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	cli = NewUserClient(conn)
	//update
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res3, err := cli.UpdateUser(ctx, 3, pb_gen.UpdateType_upwd, "", "123456")
	log.Printf("%v", res3.User)

	//delete user
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res4, err := cli.DeleteUser(ctx, 3, "guohao", "123456")
	log.Printf("%v", res4.Message)
}
