package project

import (
	"MS_Local/config"
	"MS_Local/pb_gen"
	"context"
	"io"
	"log"
	"os"
	"time"
)

func (cli *LocalClient) Download(fid string, uid uint64, fileType pb_gen.FileType) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req := &pb_gen.DownloadRequest{
		FileId:   fid,
		UserId:   uid,
		FileType: fileType,
	}
	stream, err := cli.service.Download(ctx, req)
	if err != nil {
		log.Printf("cannot download file: ", err)
		return err
	}
	rec, err := stream.Recv() //receive metadata
	if err != nil {
		log.Printf("receive metadata failed, err=[%v]", err)
		return err
	}
	metadata := rec.GetMetadata()
	fpath, err := cli.receiveStream(stream, metadata)
	if err != nil {
		return err
	}
	log.Printf("%v", *metadata)
	log.Printf(fpath)

	//err = stream.CloseSend()
	//if err!=nil{
	//	log.Printf("client close send failed, err=[%v]", err)
	//	return err
	//}
	return nil
}

func (cli *LocalClient) receiveStream(stream pb_gen.MSLocal_DownloadClient, metadata *pb_gen.DownloadMetadate) (string, error) {
	fo, err := os.CreateTemp(config.TempFilePath, "temp_download_")
	if err != nil {
		log.Printf("create temp file fail, err=[%v]", err)
		return "", err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()
	dataSize := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("reseive done, file size is %dkb", dataSize/1024)
			break
		}
		if err != nil {
			log.Printf("receive chunk failed, error=[%v]", err)
			return "", err
		}
		chunk := req.GetContent()
		dataSize += len(chunk)
		_, err = fo.Write(chunk)
		if err != nil {
			log.Printf("cannot write chunk data: err=[%v]", err)
			return "", err
		}
	}
	return fo.Name(), nil //name is path
}
