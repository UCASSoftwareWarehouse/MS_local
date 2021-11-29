package project

import (
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func GetProject(req *pb_gen.GetProjectRequest, stream pb_gen.MSLocal_GetProjectServer) error {
	var pros []model.Project

	err := project.GetProjectsByUserId(mysql.Mysql, req.Uid, int(req.Limit), int(req.Page), &pros)
	if err != nil {
		return err
	}

	for i, pro := range pros {

		res := &pb_gen.GetProjectResponse{
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
			},
		}
		err := stream.Send(res)
		if err != nil {
			log.Printf("send falied, err=[%v]", err)
			return err
		}
		log.Printf("send %d project", i)
	}
	return nil

}
