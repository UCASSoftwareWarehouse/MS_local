package mysql

import (
	"MS_Local/config"
	"testing"
)

func TestInitMysql(t *testing.T) {
	config.InitConfig()
	err := InitMysql()
	if err != nil {
		t.Errorf("%v", err)
	}
}
