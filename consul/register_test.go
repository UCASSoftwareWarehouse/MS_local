package consul

import (
	"MS_Local/config"
	"testing"
)

func TestRegister_Register(t *testing.T) {
	config.InitConfig()
	r := NewConsulRegister()
	err := r.Register()
	if err != nil {
		t.Logf("err=[%v]", err)
	}

}
