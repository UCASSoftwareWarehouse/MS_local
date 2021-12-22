package action

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"log"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()

	err := DownloadGridFile("D:\\GolangProjects\\src\\test\\test1.pdf", "61c35992657dad530833a8dc")
	if err != nil {
		t.Errorf("download error, err=[%v]", err)
	}
	//log.Printf("test")
}

func TestUploadFile(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()
	fid, err := UploadGridFile("D:\\GolangProjects\\src\\test\\test.pdf", "test.pdf")
	if err != nil {
		t.Errorf("upload error, err=[%v]", err)
	}
	log.Println(fid)
}

func TestDeleteFile(t *testing.T) {
	config.InitConfig()
	mongodb.InitMongo()
	err := DeleteGridFile("61c35992657dad530833a8dc")
	if err != nil {
		t.Errorf("delete error, err=[%v]", err)
	}
}
