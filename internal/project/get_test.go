package project

import (
	"MS_Local/client"
	"MS_Local/client/project"
	"context"
	"testing"
	"time"
)

func TestGetProject(t *testing.T) {
	conn := client.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := cli.GetProject(ctx, 1, 4, 2)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}
