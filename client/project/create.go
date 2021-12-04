package project

import (
	"MS_Local/pb_gen"
	"context"
	"log"
	"time"
)

func (cli *LocalClient) Create(projectName string, uid uint64, tags string, license string, projectDescription string, classifiers uint32) error {
	req := &pb_gen.CreateProjectRequest{
		ProjectName:        projectName,
		UserId:             uid,
		Tags:               tags,
		License:            license,
		Classifiers:        classifiers,
		ProjectDescription: projectDescription,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := cli.service.CreateProject(ctx, req)
	if err != nil {
		log.Println("create project failed, err =  [%v]", err)
		return err
	}
	log.Printf("create project success, project info is \n %v", res.ProjectInfo)
	return nil
}
