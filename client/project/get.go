package project

import (
	"MS_Local/pb_gen"
	"context"
	"io"
	"log"
)

func (cli *LocalClient) GetProject(ctx context.Context, uid uint64, page int, limit int) error {
	var pros []pb_gen.Project
	req := &pb_gen.GetProjectRequest{
		Uid:   uid,
		Limit: uint32(limit),
		Page:  uint32(page),
	}
	stream, err := cli.service.GetProject(ctx, req)

	if err != nil {
		log.Printf("can't get Project, err=[%v]", err)
		return err
	}
	count := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("receive failed, err=[%v]", err)
			return err
		}
		pros = append(pros, *res.ProjectInfo)
		log.Println(*res.ProjectInfo)
		count += 1
		log.Printf("receive %d project info", count)
	}
	log.Printf("get project info success")
	return nil
}
