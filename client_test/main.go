package main

import (
	"MS_Local/pb_gen"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err=[%v]", err)
	}
	defer conn.Close()
	cli := pb_gen.NewMSLocalClient(conn)
	//test hello
	res, err := cli.SayHello(context.Background(), &pb_gen.HelloRequest{
		Name: "Test!!!",
	})
	if err != nil {
		log.Fatal("could not greet:%v", err)
	}
	log.Printf("greeting1:%s", res.GetMessage())
	////test add user
	//res1, err := cli.RegisterUser(context.Background(), &pb_gen.RegisterUserRequest{Name: "user1",Password: "123456"})
	//if err != nil {
	//	log.Fatal("add user failed, error: ", err)
	//}
	//log.Printf("greeting:%v", res1.User)
	//test create project
	res1, err := cli.CreateProject(context.Background(), &pb_gen.CreateProjectRequest{
		ProjectName:        "project1",
		UserId:             1,
		Tags:               "v1.0",
		License:            "test_liscense",
		ProjectDescription: "this is a project",
	})
	if err != nil {
		log.Fatal("could not greet:%v", err)
	}
	log.Printf("create project:%s", res1.GetMessage())
	log.Printf("project info:\n %v", res1.ProjectInfo)

}
