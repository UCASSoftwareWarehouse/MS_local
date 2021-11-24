package binary

import (
	"MS_Local/mongodb/model"
	"MS_Local/utils/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func AddBinary(ctx context.Context, collection *mongo.Collection, binary *model.Binary) (*primitive.ObjectID, error) {
	insertResult, err := collection.InsertOne(ctx, *mongodb.Struct2Bson(*binary))
	if err != nil {
		log.Printf("add binary error, err=[%v]", err)
		return nil, err
	}
	fid := insertResult.InsertedID.(primitive.ObjectID)
	return &fid, err
}

//func AddBinarys(collection *mongo.Collection, binarys []model.Binary) ([]interface{}, error) {
//	var binarys_i []interface{}
//	for _, binary := range binarys {
//		binarys_i = append(binarys_i, mongodb.Struct2Bson(binary))
//	}
//
//	insetManyResult, err := collection.InsertMany(context.TODO(), binarys_i)
//	if err != nil {
//		log.Printf("add many binary error, err=[%v]", err)
//	}
//	return insetManyResult.InsertedIDs, err
//}

func GetBinaryByProjectId(ctx context.Context, collection *mongo.Collection, id uint64) (*model.Binary, error) {
	binary := new(model.Binary)
	filter := bson.M{model.BinaryColumns.ProjectID: id}
	err := collection.FindOne(ctx, filter).Decode(binary)
	if err != nil {
		log.Printf("get binary by file id error, err=[%v]", err)
		return nil, err
	}
	log.Println("collection.FindOne: ", binary.FileName)
	return binary, err
}

func GetBinaryByFileId(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*model.Binary, error) {
	binary := new(model.Binary)
	filter := bson.M{model.BinaryColumns.FileID: id}
	err := collection.FindOne(ctx, filter).Decode(binary)
	if err != nil {
		log.Printf("get binary by file id error, err=[%v]", err)
		return nil, err
	}
	log.Println("binary FindOne: ", binary.FileName)
	return binary, err
}

//func ReplaceOneBinary(collection *mongo.Collection, binary model.Binary) error {
//	filter := bson.M{model.BinaryColumns.FileID: binary.FileID}
//	result, err := collection.ReplaceOne(context.Background(), filter, binary)
//	if err != nil {
//		log.Printf("replace one binary file error, err=[%v]", err)
//	}
//	log.Println("collection.RepalceOne:", result)
//	return err
//}

//func UpdateOneBinaryByFileId(collection *mongo.Collection, id primitive.ObjectID, update bson.D) (mongo.UpdateResult, error) {
//	filter := bson.M{model.BinaryColumns.FileID: id}
//	result, err := collection.UpdateOne(context.Background(), filter, update)
//	if err != nil {
//		log.Printf("update one binary file error, err=[%v]", err)
//	}
//	log.Println("UpdateOne result:", result)
//	return *result, err
//}

func DeleteOneBinaryByFileId(ctx context.Context, collection mongo.Collection, id primitive.ObjectID) error {
	filter := bson.M{model.BinaryColumns.FileID: id}

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	deleteResult, err := collection.DeleteOne(ctx, filter, opts)
	if err != nil {
		log.Printf("Delete one by file id error, err=[%v]", err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
	return err
}

//func DeleteOneBinaryFile(collection *mongo.Collection, binaryfile model.Binary) error {
//	err := DeleteOneBinaryByFileId(*collection, binaryfile.FileID)
//	return err
//}
