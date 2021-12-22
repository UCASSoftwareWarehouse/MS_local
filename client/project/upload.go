package project

import (
	"MS_Local/pb_gen"
	"bufio"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

type LocalClient struct {
	service pb_gen.MSLocalClient
}

func NewLocalClient(cc *grpc.ClientConn) *LocalClient {
	cli := pb_gen.NewMSLocalClient(cc)
	return &LocalClient{cli}
}

func (cli *LocalClient) Upload(uid uint64, pid uint64, fpath string, fileType pb_gen.FileType) error {
	file, err := os.Open(fpath)
	if err != nil {
		log.Printf("cannot open file: ", err)
		return err
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := cli.service.Upload(ctx)
	if err != nil {
		log.Printf("cannot upload file: ", err)
		return err
	}
	finfo, _ := os.Stat(fpath)
	if err != nil {
		log.Printf("cannot get file info ", err)
		return err
	}
	req := &pb_gen.UploadRequest{
		Data: &pb_gen.UploadRequest_Metadata{
			Metadata: &pb_gen.UploadMetadata{
				ProjectId: pid,
				UserId:    uid,
				//ProjectName: ,
				FileInfo: &pb_gen.FileInfo{
					FileName:   finfo.Name(),
					FileType: fileType,
				},

			},
		},
	}
	log.Printf("request is:\n%v", req)
	err = stream.Send(req)
	if err != nil {
		log.Printf("cannot send image info to server: ", err, stream.RecvMsg(nil))
		return err
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot read chunk to buffer: ", err)
			return err
		}
		req := &pb_gen.UploadRequest{
			Data: &pb_gen.UploadRequest_Content{
				Content: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Printf("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return err
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("cannot receive response: ", err)
		return err
	}
	log.Printf("file info:\n%v", res.ProjectInfo)
	return nil
}
