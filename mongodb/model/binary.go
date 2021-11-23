package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Binary struct {
	FileID    primitive.ObjectID `json:"file_id" bson:"_id,omitempty"`
	FileName  string             `json:"file_name" bson:"file_name"`
	ProjectID uint64             `json:"project_id" bson:"project_id"`
	//FileType   int                  `json:"file_type" bson:"file_type"`
	FileSize   uint64              `json:"file_size" bson:"file_size"`
	UpdateTime primitive.Timestamp `json:"update_time" bson:"update_time"`
	Content    []byte              `json:"content,omitempty" bson:"content,omitempty"`
	//ChildFiles []primitive.ObjectID `json:"child_files,omitempty" bson:"child_files,omitempty"`
}

// TableName get sql table name.获取数据库表名
func (m *Binary) TableName() string {
	return "Binary"
}

// get codefile column name in mongodb.获取数据库列名
var BinaryColumns = struct {
	FileID    string
	FileName  string
	ProjectID string
	//FileType   string
	FileSize   string
	Content    string
	UpdateTime string
	//ChildFiles string
}{
	FileID:    "_id",
	FileName:  "file_name",
	ProjectID: "project_id",
	//FileType:   "file_type",
	FileSize:   "file_size",
	Content:    "content",
	UpdateTime: "update_time",
	//ChildFiles: "child_files",
}
