package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigurationEnv string

const (
	DevEnv ConfigurationEnv = "dev"
	PrdEnv ConfigurationEnv = "prd"
)

type Configuration map[ConfigurationEnv]*EachConfig

type EachConfig struct {
	AppName       string `yaml:"app_name"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	MongodbAddr   string `yaml:"mongodb_addr"`
	MongodbUser   string `yaml:"mongodb_user"`
	MongodbPwd    string `yaml:"mongodb_pwd"`
	MongodbDb     string `yaml:"mongodb_db"`
	MongodbBinary string `yaml:"mongodb_binary"`
	MongodbCode   string `yaml:"mongodb_code"`
	MysqlAddr     string `yaml:"mysql_addr"`
	MysqlUser     string `yaml:"mysql_user"`
	MysqlPwd      string `yaml:"mysql_pwd"`
	MysqlDb       string `yaml:"mysql_db"`
}

const (
	//DefaultConfigFilepath = "./config.yml"
	DefaultConfigFilepath = "D:\\GolangProjects\\src\\MS_Local\\config.yml"
)

func Parse(configFilepath string) Configuration {
	println()
	if configFilepath == "" {
		configFilepath = DefaultConfigFilepath
	}
	bs, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		log.Printf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	conf := make(Configuration)
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		log.Printf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}
	return conf
}
