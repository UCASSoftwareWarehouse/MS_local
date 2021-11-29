package project

import (
	"MS_Local/client"
	"MS_Local/client/project"
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"MS_Local/pb_gen"
	"log"
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
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "619e1d50c128e0064c70a197"

	fpath, _, err := DownloadBinary(fid, "")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadCode(t *testing.T) {
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "619cf73347fb21d93ef85390"

	fpath, _, err := DownloadCode(fid, "")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadProject(t *testing.T) {
	mongodb.InitMongo()
	mysql.InitMysql()

	fpath, _, err := DownloadProject("1")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadCodes(t *testing.T) {
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "619cf73347fb21d93ef85390"
	fpath, _, err := DownloadCodes(fid, "")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}
