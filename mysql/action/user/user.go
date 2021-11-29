package user

import (
	"MS_Local/mysql/model"
	"gorm.io/gorm"
	"log"
)

func AddUser(db *gorm.DB, user model.User) (uint64, error) {
	err := db.Create(&user).Error
	if err != nil {
		log.Printf("add user failed, err:=[%v]", err)
	}
	return user.ID, err
}

func GetAllUser(db *gorm.DB) ([]model.User, error) {
	var userList []model.User
	result := db.Find(&userList)
	if result.Error != nil {
		log.Printf("get all user info failed, err:=[%v]", result.Error)
	}
	return userList, result.Error
}

func GetUserByUserId(db *gorm.DB, userId uint64) (model.User, error) {
	user := new(model.User)
	err := db.Where("id = ?", userId).First(user).Error
	if err != nil {
		log.Printf("get user by user id error, err=[%v]", err)
	}
	return *user, err
}

func GetUserByUserName(db *gorm.DB, name string) (*model.User, error) {
	user := new(model.User)
	err := db.Where("user_name = ?", name).First(user).Error
	if err != nil {
		log.Printf("get user by user id error, err=[%v]", err)
		return nil, err
	}
	return user, err
}

func UpdateUserInfo(db *gorm.DB, uid uint64, column string, value string) error {
	user, err := GetUserByUserId(db, uid)
	if err != nil {
		log.Printf("get user in update user info failed, err=[%v]", err)
		return err
	}
	err = db.Model(&user).Update(column, value).Error
	if err != nil {
		log.Printf("update user name error, err=[%v]", err)
		return err
	}
	log.Printf("change user %s to %s", column, value)
	return nil
}

//func UpdateUserName(db *gorm.DB, userId uint64, newName string) (error) {
//	//user, _ := GetUserByUserId(db, userId)
//	//err := db.Model(&user).Update("user_name", newName).Error
//	err := db.Where("id = ?", userId).Update("user_name", newName).Error
//	if err != nil {
//		log.Printf("update user name error, err=[%v]", err)
//	}
//	log.Printf("change user name to %s", newName)
//	return err
//}
//
//func UpdateUserPassword(db *gorm.DB, userId uint64, newPassword string) ( error) {
//	user, _ := GetUserByUserId(db, userId)
//	err := db.Model(&user).Update("password", newPassword).Error
//	if err != nil {
//		log.Printf("update user password error, err=[%v]", err)
//	}
//	return  err
//}

func DeleteByUserId(db *gorm.DB, id uint64) error {
	err := db.Where(&model.User{ID: id}).Delete(model.User{}).Error
	if err != nil {
		log.Printf("delete by user id error, err=[%v]", err)
		return err
	}
	log.Printf("delete user success")
	return nil
}
