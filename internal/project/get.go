package project

import (
	"MS_Local/mongodb"
	"MS_Local/mongodb/action/code"
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	mongodb2 "MS_Local/utils/mongodb"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func GetUserProjects(req *pb_gen.GetUserProjectsRequest, stream pb_gen.MSLocal_GetUserProjectsServer) error {
	log.Println("GetUserProjects: get user's all project started")
	var pros []model.Project

	err := project.GetProjectsByUserId(mysql.Mysql, req.Uid, int(req.Limit), int(req.Page), &pros)
	if err != nil {
		return err
	}

	for i, pro := range pros {

		res := &pb_gen.GetUserProjectsResponse{
			ProjectInfo: &pb_gen.Project{
				Id:          pro.ID,
				ProjectName: pro.ProjectName,
				UserId:      pro.UserID,
				Tags:        pro.Tags,
				License:     pro.Tags,
				Updatetime: &timestamppb.Timestamp{
					Seconds: pro.UpdateTime.Unix(),
					Nanos:   0,
				},
				ProjectDescription: pro.ProjectDescription,
				CodeAddr:           pro.CodeAddr,
				BinaryAddr:         pro.BinaryAddr,
				Classifiers:        utils.GetClassifier(pro.OperatingSystem, pro.ProgrammingLanguage, pro.NaturalLanguage, pro.Topic),
			},
		}
		err := stream.Send(res)
		if err != nil {
			log.Printf("send falied, err=[%v]", err)
			return err
		}
		log.Printf("send %d project", i)
	}
	log.Println("GetUserProjects: get user's all project finish")
	return nil

}

func GetProject(ctx context.Context, req *pb_gen.GetProjectRequest) (*pb_gen.GetProjectResponse, error) {
	log.Println("GetProject: get project info started")
	pid := req.Pid
	pro, err := project.GetProjectById(mysql.Mysql, pid)
	if err != nil {
		return nil, err
	}
	log.Println("GetProject: get project info finished")
	return &pb_gen.GetProjectResponse{
		ProjectInfo: &pb_gen.Project{
			Id:                 pid,
			ProjectName:        pro.ProjectName,
			UserId:             pro.UserID,
			Tags:               pro.Tags,
			License:            pro.License,
			Updatetime:         utils.Time2Timestamp(pro.UpdateTime),
			ProjectDescription: pro.ProjectDescription,
			CodeAddr:           pro.CodeAddr,
			BinaryAddr:         pro.BinaryAddr,
			Classifiers:        utils.GetClassifier(pro.OperatingSystem, pro.ProgrammingLanguage, pro.NaturalLanguage, pro.Topic),
		},
	}, nil
}

func GetCodes(req *pb_gen.GetCodesRequest, stream pb_gen.MSLocal_GetCodesServer) error {
	log.Println("GetCodes: get codes file list start")
	finfo, err := code.GetCodeByFileId(context.Background(), mongodb.CodeCol, mongodb2.String2ObjectId(req.Fid))
	if err != nil {
		return err
	}
	if finfo.FileType == 0 { //list dir file
		for _, fid := range finfo.ChildFiles {
			cinfo, err := code.GetCodeByFileId(context.Background(), mongodb.CodeCol, fid)
			if err != nil {
				return err
			}
			res := &pb_gen.GetCodesResponse{
				FileInfo: &pb_gen.FileInfo{
					FileName: cinfo.FileName,
					FileType: pb_gen.FileType_code_file,
				},
				Fid: mongodb2.ObjectId2String(fid),
			}
			if cinfo.FileType == 0 {
				res.FileInfo.FileType = pb_gen.FileType_code_dir
			}
			err = stream.Send(res)
			if err != nil {
				log.Printf("get codes, send failed, err=[%v]", err)
				return err
			}
		}
	}

	log.Println("GetCodes: get codes file list finish")
	return nil

}
