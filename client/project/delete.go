package project

import (
	"MS_Local/pb_gen"
	"context"
	"log"
	"time"
)

func (cli *LocalClient) Delete(pid uint64, uid uint64, fileType pb_gen.FileType) error {
	req := &pb_gen.DeleteProjectRequest{
		Pid:      pid,
		Uid:      uid,
		FileType: fileType,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := cli.service.DeleteProject(ctx, req)
	if err != nil {
		log.Println("delete project failed, err=[%v]", err)
		return err
	}
	log.Printf("delete project success, project info is %v", res.Message)
	return nil
}
