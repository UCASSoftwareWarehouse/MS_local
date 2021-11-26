package project

import (
	"MS_Local/mongodb"
	"MS_Local/mongodb/action/binary"
	code2 "MS_Local/mongodb/action/code"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Delete(ctx context.Context, req *pb_gen.DeleteProjectRequest) (*pb_gen.DeleteProjectResponse, error) {
	//get
	pro, err := project.GetProjectById(mysql.Mysql, req.GetPid())
	if err != nil {
		return &pb_gen.DeleteProjectResponse{
			Status:  pb_gen.ResponseStatus_fail,
			Message: "find pro failed",
		}, err
	}

	if req.FileType == pb_gen.FileType_project {
		err := DeleteProject(pro)
		if err != nil {
			return &pb_gen.DeleteProjectResponse{
				Status:  pb_gen.ResponseStatus_fail,
				Message: "delete project failed",
			}, err
		}
	} else if req.FileType == pb_gen.FileType_binary {
		err := DeleteBinary(pro.ID)
		if err != nil {
			return &pb_gen.DeleteProjectResponse{
				Status:  pb_gen.ResponseStatus_fail,
				Message: "delete binary failed",
			}, err
		}
	} else if req.FileType == pb_gen.FileType_codes {
		err := DeleteCodes(pro.ID)
		if err != nil {
			return &pb_gen.DeleteProjectResponse{
				Status:  pb_gen.ResponseStatus_fail,
				Message: "delete codes failed",
			}, err
		}
	} else {
		return &pb_gen.DeleteProjectResponse{
			Status:  pb_gen.ResponseStatus_fail,
			Message: "file type error",
		}, status.Errorf(codes.InvalidArgument, "can only delete pro/codes/binary")
	}
	//if ctx.Err() == context.Canceled {
	//	log.Print("request is canceled")
	//	return nil, status.Error(codes.Canceled, "request is canceled")
	//}
	//
	//if ctx.Err() == context.DeadlineExceeded {
	//	log.Print("deadline is exceeded")
	//	return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	//}

	return &pb_gen.DeleteProjectResponse{
		Status:  pb_gen.ResponseStatus_ok,
		Message: "success",
	}, nil
}

func DeleteProject(pro *model.Project) error {
	//delete binary
	baddr := pro.BinaryAddr
	if baddr != "" {
		err := DeleteBinary(pro.ID)
		if err != nil {
			return err
		}
	}
	//delete codes
	caddr := pro.CodeAddr
	if caddr != "" {
		err := code2.DeleteManyCodesByProjectId(context.Background(), mongodb.CodeCol, pro.ID)
		if err != nil {
			return err
		}
	}
	//delete project
	err := project.DeleteProjectById(mysql.Mysql, pro.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBinary(pid uint64) error {
	//delete
	err := binary.DeleteBinaryByProjectId(context.Background(), mongodb.BinaryCol, pid)
	if err != nil {
		return err
	}
	//update
	err = project.UpdateProject(mysql.Mysql, pid, map[string]interface{}{model.ProjectColumns.BinaryAddr: ""})
	if err != nil {
		return err
	}
	return nil
}

func DeleteCodes(pid uint64) error {
	//delete
	err := code2.DeleteManyCodesByProjectId(context.Background(), mongodb.CodeCol, pid)
	if err != nil {
		return err
	}
	//update
	err = project.UpdateProject(mysql.Mysql, pid, map[string]interface{}{model.ProjectColumns.CodeAddr: ""})
	if err != nil {
		return err
	}
	return nil
}