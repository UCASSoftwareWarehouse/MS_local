package project

import (
	"MS_Local/client"
	"MS_Local/client/project"
	"MS_Local/pb_gen"
	"testing"
)

func TestDelete(t *testing.T) {
	conn := client.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	err := cli.Delete(7, 1, pb_gen.FileType_project)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}
