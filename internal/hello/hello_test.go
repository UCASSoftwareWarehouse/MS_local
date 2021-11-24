package hello

import (
	"MS_Local/pb_gen"
	"MS_Local/server"
	"context"
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {
	conn := server.InitMSLocalClient()
	cli := pb_gen.NewMSLocalClient(conn)
	defer conn.Close()

	//test add user
	res, err := cli.SayHello(context.Background(), &pb_gen.HelloRequest{Name: "test name"})
	if err != nil {
		t.Errorf("add user failed, error= [%v]", err)
	}
	log.Printf("greeting:%v", res.Message)
}
