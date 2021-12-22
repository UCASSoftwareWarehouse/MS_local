package project

import (
	"MS_Local/client"
	"MS_Local/pb_gen"
	"context"
	"log"
	"testing"
)

func TestCreateProject(t *testing.T) {
	conn := client.InitMSLocalClient()
	cli := pb_gen.NewMSLocalClient(conn)
	defer conn.Close()

	//test create project
	res1, err := cli.CreateProject(context.Background(), &pb_gen.CreateProjectRequest{
		ProjectName:        "project1",
		UserId:             1,
		Tags:               "v1.0",
		License:            "liscense",
		Classifiers:        0,
		ProjectDescription: "this is a project",
	})
	if err != nil {
		t.Errorf("could not greet:%v", err)
	}
	log.Printf("project info:\n %v", res1.ProjectInfo)
}
