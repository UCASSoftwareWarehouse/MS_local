package project

import (
	"MS_Local/mysql"
	"MS_Local/mysql/action/project"
	"MS_Local/mysql/model"
	"MS_Local/pb_gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func SearchProject(req *pb_gen.SearchProjectRequest, stream pb_gen.MSLocal_SearchProjectServer) error {
	log.Println("SearchProject: start")
	var pros []model.Project
	if req.KeyWord == "" {
		err := project.GetLimitProjects(mysql.Mysql, int(req.Limit), int(req.Page), &pros)
		if err != nil {
			return err
		}
	} else {
		err := project.SearchProjectByName(mysql.Mysql, req.KeyWord, int(req.Limit), int(req.Page), req.Classifiers,
			&pros)
		if err != nil {
			return err
		}
	}
	for i, pro := range pros {
		res := &pb_gen.SearchProjectResponse{
			ProjectInfo: &pb_gen.Project{
				Id:          pro.ID,
				ProjectName: pro.ProjectName,
				UserId:      pro.UserID,
				Tags:        pro.Tags,
				License:     pro.License,
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
	log.Println("SearchProject: finish")
	return nil
}
