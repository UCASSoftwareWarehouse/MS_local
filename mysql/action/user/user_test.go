package user

import (
	"MS_Local/mysql"
	"MS_Local/mysql/model"
	"MS_Local/utils"
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	db, _ := mysql.InitMysql()
	var password string = "11111111"
	user := model.User{
		UserName: "user1",
		Password: utils.Encrypt(password),
	}
	id, err := AddUser(db, user)
	if err != nil {
		t.Errorf("add user error: #{err}")
	}
	fmt.Println(id)
}

func TestUpdateUserName(t *testing.T) {
	db, _ := mysql.InitMysql()
	user, err := UpdateUserName(db, 2, "user")
	if err != nil {
		t.Errorf("update user name error: #{err}")
	}
	fmt.Println(user)
}

func TestUpdateUserPassword(t *testing.T) {
	db, _ := mysql.InitMysql()
	user, err := UpdateUserPassword(db, 2, utils.Encrypt("123456"))
	if err != nil {
		t.Errorf("update user password error: #{err}")
	}
	fmt.Println(user)
}

func TestGetUserByUserId(t *testing.T) {
	db, _ := mysql.InitMysql()
	user, err := GetUserByUserId(db, 2)
	if err != nil {
		t.Errorf("get user error: #{err}")
	}
	fmt.Println(user)
}

func TestGetAllUser(t *testing.T) {
	db, _ := mysql.InitMysql()
	users, err := GetAllUser(db)
	if err != nil {
		t.Errorf("get all user error: #{err}")
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestDeleteByUserId(t *testing.T) {
	db, _ := mysql.InitMysql()
	err := DeleteByUserId(db, 2)
	if err != nil {
		t.Errorf("delete user error: #{err}")
	}
}
