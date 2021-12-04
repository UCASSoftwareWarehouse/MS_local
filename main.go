package main

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"MS_Local/server"
	"log"
)

func main() {
	config.InitConfig()
	err := mysql.InitMysql()
	if err != nil {
		log.Fatalf("init mysql error:[%v]", err)
	}
	err = mongodb.InitMongo()
	if err != nil {
		log.Fatalf("init mongodb error:[%v]", err)
	}
	log.Printf("init database success!")
	server.StartServe()
}
