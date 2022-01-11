package project

import (
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

//是否查重
func CreateProject(ctx context.Context, req *pb_gen.CreateProjectRequest) (*pb_gen.CreateProjectResponse, error) {
	log.Println("CreateProject: start create a project")
	if len(req.ProjectName) == 0 {
		log.Printf("project name must be set")
		return nil, status.Errorf(codes.InvalidArgument, "project name can not be empty")
	}

	tmp_project := model.Project{
		ProjectName:         req.ProjectName,
		UserID:              req.UserId,
		Tags:                req.Tags,
		License:             req.License,
		UpdateTime:          time.Now(),
		ProjectDescription:  req.ProjectDescription,
		OperatingSystem:     utils.GetOSValue(req.Classifiers),
		ProgrammingLanguage: utils.GetPLValue(req.Classifiers),
		NaturalLanguage:     utils.GetNLValue(req.Classifiers),
		Topic:               utils.GetToVaule(req.Classifiers),
	}

	pid, err := project.AddProject(mysql.Mysql, tmp_project)
	if err != nil {
		return nil, err
	}
	log.Printf("CreateProject: create project %s for user %d success!", req.ProjectName, req.UserId)

	return &pb_gen.CreateProjectResponse{
		ProjectInfo: &pb_gen.Project{
			Id:                 pid,
			ProjectName:        req.ProjectName,
			UserId:             req.UserId,
			Tags:               req.Tags,
			License:            req.License,
			Updatetime:         utils.Time2Timestamp(tmp_project.UpdateTime),
			ProjectDescription: req.ProjectDescription,
			Classifiers:        req.Classifiers,
			CodeAddr:           "",
			BinaryAddr:         "",
		},
	}, nil

}
