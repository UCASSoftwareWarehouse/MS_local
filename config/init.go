package config

import (
	"log"
	"runtime"
)

var Conf *EachConfig

func IsProd() bool {
	log.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

func InitConfig() {
	InitConfigDefault()
}

func InitConfigDefault() {
	//相对路径
	c := Parse(DefaultConfigFilepath)
	//判断mac or linux
	if IsProd() {
		Conf = c[PrdEnv] //linux
	} else {
		Conf = c[DevEnv]
	}
	log.Printf("InitConfigDefault %+v", Conf)
}
