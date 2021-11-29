package config

import "testing"

func TestInitConfig(t *testing.T) {
	InitConfig()
	t.Logf("config info:\n[%v]", Conf)
}
