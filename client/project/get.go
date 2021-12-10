package project

import (
	"MS_Local/pb_gen"
	"context"
	"io"
	"log"
)

func (cli *LocalClient) GetProject(ctx context.Context, uid uint64, page int, limit int) error {
	var pros []pb_gen.Project
	req := &pb_gen.GetUserProjectsRequest{
		Uid:   uid,
		Limit: uint32(limit),
		Page:  uint32(page),
	}
	stream, err := cli.service.GetUserProjects(ctx, req)

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
func (cli *LocalClient)GetCodes(ctx context.Context, pid uint64, uid uint64,  fid string, page int, limit int)error{
	var codes []pb_gen.FileInfo
	req := &pb_gen.GetCodesRequest{
		Pid:   pid,
		Uid:   uid,
		Fid:   fid,
		Page: uint32(page),
		Limit: uint32(limit),
	}
	 stream, err := cli.service.GetCodes(ctx, req)
	 if err!=nil{
		 log.Printf("can't get the content of codes, err=[%v]",err)
	 }
	for{
		res, err:=stream.Recv()
		if err == io.EOF{
			break;
		}
		if err!=nil{
			log.Printf("receive failed, err=[%v]", err)
			return err
		}
		codes = append(codes, *res.FileInfo)
	}
	for _, code:=range(codes){
		log.Println(code)
	}
	return nil
}