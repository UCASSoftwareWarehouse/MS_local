package User

import (
	"MS_Local/pb_gen"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err=[%v]", err)
	}
	defer conn.Close()
	cli := pb_gen.NewMSLocalClient(conn)

	//test add user
	res1, err := cli.RegisterUser(context.Background(), &pb_gen.RegisterUserRequest{Name: "guohao", Password: "123456"})
	if err != nil {
		t.Errorf("add user failed, error= [%v]", err)
	}
	log.Printf("greeting:%v", res1.User)
}
