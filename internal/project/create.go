package project

import (
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

//是否查重
func CreateProject(ctx context.Context, req *pb_gen.CreateProjectRequest) (*pb_gen.CreateProjectResponse, error) {
	if len(req.ProjectName) == 0 {
		log.Printf("project name must be set")
		return &pb_gen.CreateProjectResponse{
			ProjectInfo: nil,
			Status:      pb_gen.ResponseStatus_fail,
			Message:     "project name can not be empty",
		}, status.Errorf(codes.InvalidArgument, "project name can not be empty")
	}

	tmp_project := model.Project{
		ProjectName:        req.ProjectName,
		UserID:             req.UserId,
		Tags:               req.Tags,
		License:            req.License,
		UpdateTime:         time.Now(),
		ProjectDescription: req.ProjectDescription,
	}

	pid, err := project.AddProject(mysql.Mysql, tmp_project)
	if err != nil {
		return &pb_gen.CreateProjectResponse{
			ProjectInfo: nil,
			Status:      pb_gen.ResponseStatus_fail,
			Message:     fmt.Sprintf("add user failed, err=%v", err),
		}, status.Errorf(codes.Internal, fmt.Sprintf("add user failed, err=%v", err))
	}
	log.Printf("create project %s for user %d success!", req.ProjectName, req.UserId)

	return &pb_gen.CreateProjectResponse{
		ProjectInfo: &pb_gen.Project{
			Id:                 pid,
			ProjectName:        req.ProjectName,
			UserId:             req.UserId,
			Tags:               req.Tags,
			License:            req.License,
			Updatetime:         utils.Time2Timestamp(tmp_project.UpdateTime),
			ProjectDescription: req.ProjectDescription,
			CodeAddr:           "",
			BinaryAddr:         "",
		},
		Status:  pb_gen.ResponseStatus_ok,
		Message: "create project success!",
	}, nil

}
