package model

import "time"

// Project [...]
type Project struct {
	ID                  uint64    `gorm:"primaryKey;column:id;type:bigint unsigned;not null" json:"-"`
	ProjectName         string    `gorm:"index:project_name,class:FULLTEXT;column:project_name;type:varchar(100);not null" json:"projectName"`
	UserID              uint64    `gorm:"index:user_id;column:user_id;type:bigint unsigned;not null" json:"userId"`
	User                User      `gorm:"joinForeignKey:user_id;foreignKey:id" json:"userList"`
	OperatingSystem     uint8     `gorm:"column:operating_system;type:tinyint unsigned;not null" json:"operatingSystem"`
	ProgrammingLanguage uint8     `gorm:"column:programming_language;type:tinyint unsigned;not null" json:"programmingLanguage"`
	NaturalLanguage     uint8     `gorm:"column:natural_language;type:tinyint unsigned;not null" json:"naturalLanguage"`
	Topic               uint8     `gorm:"column:topic;type:tinyint unsigned;not null" json:"topic"`
	Tags                string    `gorm:"column:tags;type:varchar(20)" json:"tags"`
	CodeAddr            string    `gorm:"column:code_addr;type:char(24)" json:"codeAddr"`
	BinaryAddr          string    `gorm:"column:binary_addr;type:char(24)" json:"binaryAddr"`
	License             string    `gorm:"column:license;type:varchar(50)" json:"license"`
	UpdateTime          time.Time `gorm:"column:update_time;type:timestamp;not null" json:"updateTime"`
	ProjectDescription  string    `gorm:"column:project_description;type:text" json:"projectDescription"`
}

// TableName get sql table name.获取数据库表名
func (m *Project) TableName() string {
	return "project"
}

// ProjectColumns get sql column name.获取数据库列名
var ProjectColumns = struct {
	ID                  string
	ProjectName         string
	UserID              string
	OperatingSystem     string
	ProgrammingLanguage string
	NaturalLanguage     string
	Topic               string
	Tags                string
	CodeAddr            string
	BinaryAddr          string
	License             string
	UpdateTime          string
	ProjectDescription  string
}{
	ID:                  "id",
	ProjectName:         "project_name",
	UserID:              "user_id",
	OperatingSystem:     "operating_system",
	ProgrammingLanguage: "programming_language",
	NaturalLanguage:     "natural_language",
	Topic:               "topic",
	Tags:                "tags",
	CodeAddr:            "code_addr",
	BinaryAddr:          "binary_addr",
	License:             "license",
	UpdateTime:          "update_time",
	ProjectDescription:  "project_description",
}
