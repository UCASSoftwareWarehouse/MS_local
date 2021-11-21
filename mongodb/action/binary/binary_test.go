package binary

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mongodb/model"
	mongodb2 "MS_Local/utils/mongodb"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
	"time"
)

func TestAddBinary(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)

	binary := model.Binary{
		//FileID: primitive.NewObjectID(),
		FileName:   "test",
		ProjectID:  1,
		FileSize:   16,
		UpdateTime: mongodb2.Time2Timestamp(time.Now()),
		Content:    []byte("hello mongodb"),
	}
	id, err := AddBinary(collection, binary)
	if err != nil {
		t.Errorf("add code file error: %s", err)
	}
	fmt.Println("insert id is ", id)
}

func TestAddBinarys(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var binaryfiles = []model.Binary{
		model.Binary{
			//FileID: primitive.ObjectID{},
			FileName:  "file1",
			ProjectID: 1,
			//FileType: 1,
			FileSize:   16,
			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
			Content:    []byte("this is file one"),
		},
		model.Binary{
			//FileID: primitive.ObjectID{},
			FileName:  "file2",
			ProjectID: 1,
			//FileType: 1,
			FileSize:   16,
			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
			Content:    []byte("this is file two"),
		},
		model.Binary{
			//FileID: primitive.ObjectID{},
			FileName:  "file3",
			ProjectID: 1,
			//FileType: 1,
			FileSize:   16,
			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
			Content:    []byte("this is file three"),
		},
	}
	ids, err := AddBinarys(collection, binaryfiles)
	if err != nil {
		t.Errorf("add binaryfiles failed: #{err}")
	}
	fmt.Println(ids)
}

func TestGetBinaryByProjectId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	binarys, err := GetBinaryByProjectId(collection, 1)
	if err != nil {
		t.Errorf("test get binarys by projectid failded")
	}
	for _, binaryfile := range binarys {
		fmt.Println(reflect.TypeOf(binaryfile))
	}
}

func TestGetBinaryByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var stringid = "6199ff8a221e83dc617760bb"
	binary, err := GetBinaryByFileId(collection, mongodb2.String2ObjectId(stringid))
	if err != nil {
		t.Errorf("test get codefile by file id")
	}
	fmt.Println(reflect.TypeOf(binary))
	fmt.Println(binary)
}

func TestReplaceOneBinary(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var stringid = "6199ff8a221e83dc617760bb"
	binary, _ := GetBinaryByFileId(collection, mongodb2.String2ObjectId(stringid))
	binary.FileSize = 20
	binary.UpdateTime = mongodb2.Time2Timestamp(time.Now())
	err := ReplaceOneBinary(collection, binary)
	if err != nil {
		t.Errorf("test replace binary failed")
	}
}

func TestUpdateOneBinaryByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var stringid = "6199ff8a221e83dc617760bb"
	update := bson.D{{"$set", bson.D{{model.BinaryColumns.FileSize, 1}}}}
	_, err := UpdateOneBinaryByFileId(collection, mongodb2.String2ObjectId(stringid), update)
	if err != nil {
		t.Errorf("test update codefile failed")
	}
}

func TestDeleteOneBinaryFile(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var stringid = "6199ff8a221e83dc617760bb"
	binaryfile, _ := GetBinaryByFileId(collection, mongodb2.String2ObjectId(stringid))
	err := DeleteOneBinaryFile(collection, binaryfile)
	if err != nil {
		t.Errorf("test delete one codefile failed")
	}
}
