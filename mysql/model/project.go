package model

import (
	"time"
)

// https://github.com/xxjwxc/gormt/blob/master/README_zh_cn.md
// start gormt.exe
// Project [...]
//自动级联删除，不知道为什么
type Project struct {
	ID          uint64 `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"-"`
	ProjectName string `gorm:"column:project_name;type:varchar(100);not null" json:"projectName"`
	UserID      uint64 `gorm:"index:user_id;column:user_id;type:bigint unsigned;not null" json:"userId"`
	//User               User      `gorm:"joinForeignKey:user_id;foreignKey:id" json:"userList"`
	Tags               string    `gorm:"column:tags;type:varchar(20)" json:"tags,omitempty"`
	CodeAddr           string    `gorm:"column:code_addr;type:char(12)" json:"codeAddr,omitempty"`
	BinaryAddr         string    `gorm:"column:binary_addr;type:char(12)" json:"binaryAddr"`
	License            string    `gorm:"column:license;type:varchar(50)" json:"license"`
	UpdateTime         time.Time `gorm:"column:update_time;type:timestamp;not null" json:"updateTime"`
	ProjectDescription string    `gorm:"column:project_description;type:text" json:"projectDescription,omitempty"`
}

// TableName get sql table name.获取数据库表名
func (m *Project) TableName() string {
	return "project"
}

// ProjectColumns get sql column name.获取数据库列名
var ProjectColumns = struct {
	ID                 string
	ProjectName        string
	UserID             string
	Tags               string
	CodeAddr           string
	BinaryAddr         string
	License            string
	UpdateTime         string
	ProjectDescription string
}{
	ID:                 "id",
	ProjectName:        "project_name",
	UserID:             "user_id",
	Tags:               "tags",
	CodeAddr:           "code_addr",
	BinaryAddr:         "binary_addr",
	License:            "license",
	UpdateTime:         "update_time",
	ProjectDescription: "project_description",
}
