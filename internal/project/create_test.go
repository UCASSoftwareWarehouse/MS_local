package project

import (
	"MS_Local/pb_gen"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestCreateProject(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err=[%v]", err)
	}
	defer conn.Close()
	cli := pb_gen.NewMSLocalClient(conn)

	//test create project
	res1, err := cli.CreateProject(context.Background(), &pb_gen.CreateProjectRequest{
		ProjectName:        "project1",
		UserId:             2,
		Tags:               "v1.0",
		License:            "test_liscense",
		ProjectDescription: "this is a project",
	})
	if err != nil {
		t.Errorf("could not greet:%v", err)
	}
	log.Printf("create project:%s", res1.GetMessage())
	log.Printf("project info:\n %v", res1.ProjectInfo)
}
