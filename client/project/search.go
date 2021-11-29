package project

import (
	"MS_Local/pb_gen"
	"context"
	"io"
	"log"
	"time"
)

func (cli *LocalClient) SearchProject(keyWord string, page int, limit int) error {
	var pros []pb_gen.Project
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req := &pb_gen.SearchProjectRequest{
		Page:    uint32(page),
		Limit:   uint32(limit),
		KeyWord: keyWord,
	}
	stream, err := cli.service.SearchProject(ctx, req)
	if err != nil {
		log.Printf("cann't search %s", keyWord)
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
	return nil
}
