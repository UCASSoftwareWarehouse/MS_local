package mysql

import (
	"MS_Local/config"
	"fmt"
	"gorm.io/driver/mysql"
	"log"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mysql *gorm.DB

// InitMysql 数据库连接
func InitMysql() error {
	log.Printf("init mysql....")
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
	//	config2.MysqlUser, config2.MysqlPassword, config2.MysqlHost, config2.MysqlPort, config2.MysqlDatabase, config2.MysqlCharset)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
		config.Conf.MysqlUser, config.Conf.MysqlPwd, config.Conf.MysqlAddr, config.Conf.MysqlDb, "utf8")

	config := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // 字符串字段的默认大小
		DisableDatetimePrecision:  true,  // 禁用日期时间精度，MySQL 5.6 之前不支持
		DontSupportRenameIndex:    true,  // 重命名索引时删除和创建，MySQL 5.7 之前不支持重命名索引，MariaDB
		DontSupportRenameColumn:   true,  // `change` 重命名列，MySQL 8 之前不支持重命名列，MariaDB
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	DB, err := gorm.Open(mysql.New(config), &gorm.Config{})
	if err != nil {
		//log.Fatal("open mysql failed, err =[%s]", err)
		log.Fatal("test")
	}
	Mysql = DB

	return err
}
