package MS_Local

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mysql"
	"MS_Local/server"
)

func main() {
	config.InitConfig()
	mysql.InitMysql()
	mongodb.InitMongo()
	server.StartServe()
}
