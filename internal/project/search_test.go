package project

import (
	"MS_Local/client"
	"MS_Local/client/project"
	"testing"
)

func TestSearchProject(t *testing.T) {
	conn := client.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	err := cli.SearchProject("jec", 1, 10)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}
