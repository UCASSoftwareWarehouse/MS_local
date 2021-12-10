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

func (s *MSLocalServer) Download(req *pb_gen.DownloadRequest, stream pb_gen.MSLocal_DownloadServer) error {
	return project.Download(req, stream)
}

func (s *MSLocalServer) DeleteProject(ctx context.Context, req *pb_gen.DeleteProjectRequest) (*pb_gen.DeleteProjectResponse, error) {
	return project.Delete(ctx, req)
}

func (s *MSLocalServer) SearchProject(req *pb_gen.SearchProjectRequest, stream pb_gen.MSLocal_SearchProjectServer) error {
	return project.SearchProject(req, stream)
}

func (s *MSLocalServer) GetUserProjects(req *pb_gen.GetUserProjectsRequest, stream pb_gen.MSLocal_GetUserProjectsServer) error {
	return project.GetUserProjects(req, stream)
}
func (s *MSLocalServer) GetCodes(req *pb_gen.GetCodesRequest, stream pb_gen.MSLocal_GetCodesServer) error {
	return project.GetCodes(req, stream)
}

func (s *MSLocalServer)GetProject(ctx context.Context, req *pb_gen.GetProjectRequest)(*pb_gen.GetProjectResponse, error){
	return project.GetProject(ctx, req)
}

