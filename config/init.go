package config

import (
	"log"
)

var Conf *EachConfig

func InitConfig() {
	InitConfigDefault()
}

func InitConfigDefault() {
	//相对路径
	c := Parse("")
	//判断是否是 linux
	env := GetEnv()
	Conf = c[env]
	log.Printf("InitConfigDefault %+v", Conf)
}
