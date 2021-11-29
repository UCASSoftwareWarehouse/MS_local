package user

import (
	"MS_Local/config"
	"MS_Local/mysql"
	"MS_Local/mysql/model"
	"MS_Local/utils"
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	var password string = "123456"
	user := model.User{
		UserName: "guohao",
		Password: utils.Encrypt(password),
	}
	id, err := AddUser(mysql.Mysql, user)
	if err != nil {
		t.Errorf("add user error: #{err}")
	}
	fmt.Println(id)
}

//func TestUpdateUserName(t *testing.T) {
//	config.InitConfig()
//	mysql.InitMysql()
//	user, err := UpdateUserName(mysql.Mysql, 2, "user")
//	if err != nil {
//		t.Errorf("update user name error: #{err}")
//	}
//	fmt.Println(user)
//}
//
//func TestUpdateUserPassword(t *testing.T) {
//	db, _ := mysql.InitMysql()
//	user, err := UpdateUserPassword(db, 2, utils.Encrypt("123456"))
//	if err != nil {
//		t.Errorf("update user password error: #{err}")
//	}
//	fmt.Println(user)
//}

func TestGetUserByUserId(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	user, err := GetUserByUserId(mysql.Mysql, 2)
	if err != nil {
		t.Errorf("get user error: #{err}")
	}
	fmt.Println(user)
}

func TestGetAllUser(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	users, err := GetAllUser(mysql.Mysql)
	if err != nil {
		t.Errorf("get all user error: #{err}")
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestDeleteByUserId(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	err := DeleteByUserId(mysql.Mysql, 2)
	if err != nil {
		t.Errorf("delete user error: #{err}")
	}
}

func TestUpdateUserInfo(t *testing.T) {
	config.InitConfig()
	mysql.InitMysql()
	err := UpdateUserInfo(mysql.Mysql, 1, model.UserColumns.UserName, "admin")
	if err != nil {
		t.Errorf("update user name error, err=[%v]", err)
	}
	err = UpdateUserInfo(mysql.Mysql, 1, model.UserColumns.Password, utils.Encrypt("123456"))
	if err != nil {
		t.Errorf("update user password error, err=[%v]", err)
	}
}
