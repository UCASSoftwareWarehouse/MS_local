package user

import (
	"MS_Local/pb_gen"
	"MS_Local/server"
	"context"
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {
	conn := server.InitMSLocalClient()
	cli := pb_gen.NewMSLocalClient(conn)
	defer conn.Close()

	//test add user
	res1, err := cli.RegisterUser(context.Background(), &pb_gen.RegisterUserRequest{Name: "guohao", Password: "123456"})
	if err != nil {
		t.Errorf("say hello failed, error= [%v]", err)
	}
	log.Printf("greeting:%v", res1.User)
}
