package routes

import (
	"MS_Local/internal/project"
	"MS_Local/pb_gen"
	"context"
)

func (s *MSLocalServer) CreateProject(ctx context.Context, req *pb_gen.CreateProjectRequest) (*pb_gen.CreateProjectResponse, error) {
	return project.CreateProject(ctx, req)
}

func (s *MSLocalServer) Upload(stream pb_gen.MSLocal_UploadServer) error {
	return project.Upload(stream)
}
