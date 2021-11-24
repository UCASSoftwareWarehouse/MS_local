package project

import (
	"MS_Local/pb_gen"
	"MS_Local/server"
	"context"
	"log"
	"testing"
)

func TestCreateProject(t *testing.T) {
	conn := server.InitMSLocalClient()
	cli := pb_gen.NewMSLocalClient(conn)
	defer conn.Close()

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
