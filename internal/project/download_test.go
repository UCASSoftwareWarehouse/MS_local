package project

import (
	"MS_Local/client"
	"MS_Local/client/project"
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"MS_Local/pb_gen"
	"log"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	conn := client.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	fid := "619d2218e9f25f10df00a109"
	err := cli.Download(fid, 1, pb_gen.FileType_binary)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}

func TestDownloadBinary(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "61add42e3e24c4b998841475"

	fpath, _, err := DownloadBinary(fid, filepath.Join(config.TempFilePath, "download"))
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadCode(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "61add4653e24c4b99884148a"

	fpath, _, err := DownloadCode(fid, filepath.Join(config.TempFilePath, "download"))
	if err != nil {
		t.Error("download codes failed")
	}
	log.Printf(fpath)
}

func TestDownloadProject(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()
	mysql.InitMysql()

	fpath, _, err := DownloadProject(1)
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadCodes(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "619cf73347fb21d93ef85390"
	fpath, _, err := DownloadCodes(fid, "")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}
