package code

import (
	"MS_Local/mongodb/model"
	"MS_Local/utils/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func AddCode(ctx context.Context, collection *mongo.Collection, code *model.Code) (*primitive.ObjectID, error) {
	insertResult, err := collection.InsertOne(ctx, *mongodb.Struct2Bson(*code)) //insert为interface类型
	if err != nil {
		log.Printf("add code error, err = [%v]", err)
		return nil, err
	}
	fid := insertResult.InsertedID.(primitive.ObjectID)
	return &fid, err
}

//func AddCodes(collection *mongo.Collection, codes []model.Code) ([]interface{}, error) {
//	//转化为interface数组
//	//codes_i := make([]interface{}, len(codes))
//	var codes_i []interface{}
//	for _, code := range codes {
//		codes_i = append(codes_i, mongodb.Struct2Bson(code))
//	}
//
//	insetManyResult, err := collection.InsertMany(context.TODO(), codes_i)
//	if err != nil {
//		log.Printf("insert many error, err=[%v]", err)
//	}
//	//log.Println(insetManyResult.InsertedIDs)
//	return insetManyResult.InsertedIDs, err
//}

func GetCodeByProjectId(ctx context.Context, collection *mongo.Collection, id uint64) (*model.Code, error) {
	code := new(model.Code)
	filter := bson.M{model.CodeColumns.ProjectID: id}
	err := collection.FindOne(ctx, filter).Decode(code)
	if err != nil {
		log.Printf("get code by file id error, err=[%v]", err)
		return nil, err
	}
	log.Println("collection.FindOne: ", code.FileName)
	return code, err
}

func GetCodeByFileId(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*model.Code, error) {
	code := new(model.Code)
	filter := bson.M{model.CodeColumns.FileID: id}
	err := collection.FindOne(ctx, filter).Decode(code)
	if err != nil {
		log.Printf("get code by file id error, err=[%v]", err)
		return nil, err
	}
	//log.Println("code FindOne: ", code.FileName)
	return code, err
}

//func ReplaceOneCode(collection *mongo.Collection, code model.Code) error {
//	filter := bson.M{model.CodeColumns.FileID: code.FileID}
//	result, err := collection.ReplaceOne(context.Background(), filter, code)
//	if err != nil {
//		log.Printf("replace one code error, er=[%v]", err)
//	}
//	log.Println("collection.RepalceOne:", result)
//	return err
//}

//func UpdateOneCodeByFileId(collection *mongo.Collection, id primitive.ObjectID, update bson.D) (mongo.UpdateResult, error) {
//	filter := bson.M{model.CodeColumns.FileID: id}
//	result, err := collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		log.Printf("update one code error, er=[%v]", err)
//	}
//	log.Println("UpdateOne result:", result)
//	//codefile, _ := GetCodeByFileId(collection, id)
//	return *result, err
//}

func DeleteOneCodeByFileId(ctx context.Context, collection mongo.Collection, id primitive.ObjectID) error {
	filter := bson.M{model.CodeColumns.FileID: id}
	// Delete all documents in which the "name" field is "Bob" or "bob".
	// Specify the Collation option to provide a collation that will ignore case
	// for string comparisons.
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	deleteResult, err := collection.DeleteOne(ctx, filter, opts)
	if err != nil {
		log.Printf("delete one code error, er=[%v]", err)
		return err
	}
	log.Println("collection.DeleteOne:", deleteResult)
	return nil
}

func DeleteManyCodesByProjectId(ctx context.Context, collection *mongo.Collection, id uint64) error {
	filter := bson.M{model.CodeColumns.ProjectID: id}
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	res, err := collection.DeleteMany(ctx, filter, opts)
	if err != nil {
		log.Printf("delete many code error, er=[%v]", err)
		return err
	}
	fmt.Printf("delete %d docements\n", res.DeletedCount)
	return nil
}

//func DeleteManyCodesByFileId(collection *mongo.Collection, ids []primitive.ObjectID) {
//	for _, id := range ids {
//		err := DeleteOneCodeByFileId(*collection, id)
//		if err != nil {
//			log.Fatal(" delete file error, file id is ", id, "error is: ", err)
//		}
//	}
//}

//func UploadFile(collection *mongo.Collection, dirname string, projectid uint64) primitive.ObjectID {
//	fileinfo, _ := os.Stat(dirname)
//	codefile := model.Code{
//		FileName:   fileinfo.Name(),
//		ProjectID:  projectid,
//		FileSize:   fileinfo.Size(),
//		UpdateTime: mongodb.Time2Timestamp(fileinfo.ModTime()),
//	}
//
//	if fileinfo.IsDir() {
//		codefile.FileType = 0
//		d, err := os.Open(dirname)
//		if err != nil {
//			log.Fatal(err)
//			os.Exit(1)
//		}
//		defer d.Close()
//		files, err := d.ReadDir(-1)
//		if err != nil {
//			log.Fatal(err)
//			os.Exit(1)
//		}
//		for _, cfile := range files {
//			cid := UploadFile(collection, cfile.Name(), projectid)
//			codefile.ChildFiles = append(codefile.ChildFiles, cid)
//		}
//	} else {
//		codefile.FileType = 1
//		file, _ := os.Open(codefile.FileName)
//		defer file.Close()
//		codefile.Content, _ = ioutil.ReadAll(file)
//
//	}
//	id, _ := AddCode(collection, codefile)
//	codefile.FileID = mongodb.Interface2ObjectId(id)
//	return codefile.FileID
//}
