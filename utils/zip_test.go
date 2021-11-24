package utils

import (
	"log"
	"os"
	"testing"
)

func TestUnzip(t *testing.T) {
	//filepath.Base("\tfilepath.Base(\"\")\n")
	tempDir, err := os.MkdirTemp("D:\\GolangProjects\\src\\test", "temp")
	if err != nil {
		t.Errorf("create temp dir error, err=%v", err)
	}
	defer os.RemoveAll(tempDir)
	log.Printf(tempDir)
	//defer os.RemoveAll(tempDir)
	fpath, err := Unzip("D:\\GolangProjects\\src\\test\\MS_RemoteCode-master.zip", tempDir)
	if err != nil {
		t.Errorf("unzip failed, err=%v", err)
	}
	log.Printf("unzip file to %v", fpath)
}

func TestZip(t *testing.T) {
	fpath, err := Zip("D:\\GolangProjects\\src\\test\\MS_RemoteCode-master", "D:\\GolangProjects\\src\\test\\MS_RemoteCode-master.zip")
	if err != nil {
		t.Errorf("unzip failed, err=%v", err)
	}
	log.Printf(fpath)
}
