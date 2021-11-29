package mongodb

import (
	"MS_Local/config"
	"fmt"
	"testing"
)

func TestInitMongo(t *testing.T) {
	config.InitConfig()
	err := InitMongo()
	fmt.Println(err)
}
