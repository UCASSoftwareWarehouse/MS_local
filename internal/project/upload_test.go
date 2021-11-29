package project

import (
	"MS_Local/client"
	"MS_Local/client/project"
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"MS_Local/pb_gen"
	"MS_Local/utils"
	"os"
	"testing"
)

func TestUploader_SaveBinary(t *testing.T) {
	var fpath string = "D:\\GolangProjects\\src\\test\\gormt.exe"
	mongodb.InitMongo()
	mysql.InitMysql()
	finfo, _ := os.Stat(fpath)
	metadata := &pb_gen.UploadMetadata{
		ProjectId: 1,
		UserId:    1,
		//ProjectName: "test name",
		FileInfo: &pb_gen.FileInfo{
			FileName:   finfo.Name(),
			FileSize:   uint64(finfo.Size()),
			Updatetime: utils.Time2Timestamp(finfo.ModTime()),
		},
		FileType: pb_gen.FileType_binary,
	}
	err := uploader.SaveBinary(fpath, metadata)
	if err != nil {
		t.Errorf("save binary failed, err=[%v]", err)
	}
}

func TestUploader_SaveCodes(t *testing.T) {
	var fpath string = "D:\\GolangProjects\\src\\test\\MS_RemoteCode-master.zip"
	mongodb.InitMongo()
	mysql.InitMysql()
	finfo, _ := os.Stat(fpath)
	metadata := &pb_gen.UploadMetadata{
		ProjectId: 1,
		UserId:    1,
		//ProjectName: "test name",
		FileInfo: &pb_gen.FileInfo{
			FileName:   finfo.Name(),
			FileSize:   uint64(finfo.Size()),
			Updatetime: utils.Time2Timestamp(finfo.ModTime()),
		},
		FileType: pb_gen.FileType_codes,
	}
	err := uploader.SaveCodes(fpath, metadata)
	if err != nil {
		t.Errorf("err=[%v]", err)
	}
}

func TestUpload(t *testing.T) {
	conn := client.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	fpath := "D:\\GolangProjects\\src\\test\\gormt.exe"
	err := cli.Upload(1, 1, fpath, pb_gen.FileType_binary)
	if err != nil {
		t.Errorf("client upload failed, err=[%v]", err)
	}
}
