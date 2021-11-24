package binary

import (
	"MS_Local/config"
	"MS_Local/mongodb"
	"MS_Local/mongodb/model"
	mongodb2 "MS_Local/utils/mongodb"
	"context"
	"fmt"
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
	id, err := AddBinary(context.TODO(), collection, &binary)
	if err != nil {
		t.Errorf("add code file error: %s", err)
	}
	fmt.Println("insert id is ", id)
}

//func TestAddBinarys(t *testing.T) {
//	db, _ := mongodb.InitMongo()
//	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
//	var binaryfiles = []model.Binary{
//		model.Binary{
//			//FileID: primitive.ObjectID{},
//			FileName:  "file1",
//			ProjectID: 1,
//			//FileType: 1,
//			FileSize:   16,
//			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
//			Content:    []byte("this is file one"),
//		},
//		model.Binary{
//			//FileID: primitive.ObjectID{},
//			FileName:  "file2",
//			ProjectID: 1,
//			//FileType: 1,
//			FileSize:   16,
//			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
//			Content:    []byte("this is file two"),
//		},
//		model.Binary{
//			//FileID: primitive.ObjectID{},
//			FileName:  "file3",
//			ProjectID: 1,
//			//FileType: 1,
//			FileSize:   16,
//			UpdateTime: mongodb2.Time2Timestamp(time.Now()),
//			Content:    []byte("this is file three"),
//		},
//	}
//	ids, err := AddBinarys(collection, binaryfiles)
//	if err != nil {
//		t.Errorf("add binaryfiles failed: #{err}")
//	}
//	fmt.Println(ids)
//}

func TestGetBinaryByProjectId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	binary, err := GetBinaryByProjectId(context.Background(), collection, 7)
	if err != nil {
		t.Errorf("test get binary by projectid failded")
	}
	fmt.Println(reflect.TypeOf(binary))
}

func TestGetBinaryByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var stringid = "619d2218e9f25f10df00a109"
	binary, err := GetBinaryByFileId(context.Background(), collection, mongodb2.String2ObjectId(stringid))
	if err != nil {
		t.Errorf("test get codefile by file id")
	}
	fmt.Println(reflect.TypeOf(binary))
}

//func TestReplaceOneBinary(t *testing.T) {
//	db, _ := mongodb.InitMongo()
//	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
//	var stringid = "6199ff8a221e83dc617760bb"
//	binary, _ := GetBinaryByFileId(context.Background(),collection, mongodb2.String2ObjectId(stringid))
//	binary.FileSize = 20
//	binary.UpdateTime = mongodb2.Time2Timestamp(time.Now())
//	err := ReplaceOneBinary(collection, *binary)
//	if err != nil {
//		t.Errorf("test replace binary failed")
//	}
//}

//func TestUpdateOneBinaryByFileId(t *testing.T) {
//	db, _ := mongodb.InitMongo()
//	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
//	var stringid = "6199ff8a221e83dc617760bb"
//	update := bson.D{{"$set", bson.D{{model.BinaryColumns.FileSize, 1}}}}
//	_, err := UpdateOneBinaryByFileId(collection, mongodb2.String2ObjectId(stringid), update)
//	if err != nil {
//		t.Errorf("test update codefile failed")
//	}
//}

func TestDeleteOneBinaryByFileId(t *testing.T) {
	db, _ := mongodb.InitMongo()
	collection := mongodb.GetCollectionFromMongo(db, config.BinaryCollection)
	var stringid = "619dda3a215c2f4ed65ff317"
	err := DeleteOneBinaryByFileId(context.Background(), *collection, mongodb2.String2ObjectId(stringid))
	if err != nil {
		t.Errorf("test delete one codefile failed")
	}
}
