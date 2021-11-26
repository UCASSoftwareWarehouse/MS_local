package project

import (
	"MS_Local/client/project"
	"MS_Local/pb_gen"
	"MS_Local/server"
	"testing"
)

func TestDelete(t *testing.T) {
	conn := server.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	err := cli.Delete(7, 1, pb_gen.FileType_project)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}
