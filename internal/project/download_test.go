package project

import (
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"log"
	"testing"
)

func TestDownloadBinary(t *testing.T) {
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "619cef2ac64889f1647a2f8d"

	fpath, err := DownloadBinary(fid, "")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadCode(t *testing.T) {
	mongodb.InitMongo()
	mysql.InitMysql()
	fid := "619cf73347fb21d93ef8537c"

	fpath, err := DownloadCode(fid, "")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}

func TestDownloadProject(t *testing.T) {
	mongodb.InitMongo()
	mysql.InitMysql()

	fpath, err := DownloadProject("1")
	if err != nil {
		t.Error("download binary failed")
	}
	log.Printf(fpath)
}
