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
	err := cli.GetProject(ctx, 1, 2, 2)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}

func TestGetCodes(t *testing.T) {
	conn := client.InitMSLocalClient()
	defer conn.Close()
	cli := project.NewLocalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := cli.GetCodes(ctx, 1, 2, "61c3651140baa5895b598824", 1, 5)
	if err != nil {
		t.Errorf("client download failed, err=[%v]", err)
	}
}