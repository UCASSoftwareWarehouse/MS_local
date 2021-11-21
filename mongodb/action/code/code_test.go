package code

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mongodb/model"
	mongodb2 "MS_Local/utils/mongodb"
	"fmt"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestAddCode(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)

	code := model.Code{
		//FileID: primitive.NewObjectID(),
		FileName:   "test",
		ProjectID:  1,
		FileType:   1,
		FileSize:   16,
		UpdateTime: mongodb2.Time2Timestamp(time.Now()),
		Content:    []byte("hello mongodb"),
	}
	id, err := AddCode(collection, code)
	if err != nil {
		t.Errorf("add code file error: %s", err)
	}
	fmt.Println("insert id is ", id)
}

func TestAddCodes(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	var codes = []model.Code{
		model.Code{
			//FileID: primitive.ObjectID{},
			FileName:   "file1",
			ProjectID:  1,
			FileType:   1,
			FileSize:   16,
			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
			Content:    []byte("this is file one"),
		},
		model.Code{
			//FileID: primitive.ObjectID{},
			FileName:   "file2",
			ProjectID:  1,
			FileType:   1,
			FileSize:   16,
			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
			Content:    []byte("this is file two"),
		},
		model.Code{
			//FileID: primitive.ObjectID{},
			FileName:   "file3",
			ProjectID:  1,
			FileType:   1,
			FileSize:   16,
			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
			Content:    []byte("this is file three"),
		},
	}
	ids, err := AddCodes(collection, codes)
	if err != nil {
		t.Errorf("add codes failed: #{err}")
	}
	fmt.Println(ids)
}

func TestGetCodesByProjectId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	codes, err := GetCodesByProjectId(collection, 1)
	if err != nil {
		t.Errorf("test get codes by projectid failded")
	}
	for _, codefile := range codes {
		fmt.Println(reflect.TypeOf(codefile))
	}
}

func TestGetCodeByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	var stringid = "619a00806328cd73d84d6be0"
	code, err := GetCodeByFileId(collection, mongodb2.String2ObjectId(stringid))
	if err != nil {
		t.Errorf("test get code by file id")
	}
	fmt.Println(reflect.TypeOf(code))
	fmt.Println(code)
}

func TestReplaceOneCode(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	var stringid = "619a00806328cd73d84d6be0"
	code, _ := GetCodeByFileId(collection, mongodb2.String2ObjectId(stringid))
	code.FileSize = 20
	code.UpdateTime = mongodb2.Time2Timestamp(time.Now())
	err := ReplaceOneCode(collection, code)
	if err != nil {
		t.Errorf("test replace code failed")
	}
}

func TestUpdateOneCodeByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	var stringid = "619a00806328cd73d84d6be0"
	update := bson.D{{"$set", bson.D{{model.CodeColumns.FileSize, 1}}}}
	_, err := UpdateOneCodeByFileId(collection, mongodb2.String2ObjectId(stringid), update)
	if err != nil {
		t.Errorf("test update codefile failed")
	}
}

func TestDeleteOneCodeByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	var stringid = "619a00806328cd73d84d6be0"
	code, _ := GetCodeByFileId(collection, mongodb2.String2ObjectId(stringid))
	err := DeleteOneCode(collection, code)
	if err != nil {
		t.Errorf("test delete one code failed")
	}
}

func TestDeleteManyCodesByProjectId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.CodeCollection)
	err := DeleteManyCodesByProjectId(collection, 1)
	if err != nil {
		t.Errorf("test delete many by projectID failed")
	}
}
