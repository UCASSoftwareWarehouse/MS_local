package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func String2ObjectId(s string) (primitive.ObjectID, error) {
	objId, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Printf("string %s to object id error, err=[%v]",s, err)
		return primitive.NilObjectID, err
	}
	return objId,nil
}

func ObjectId2String(id primitive.ObjectID) string {
	string := id.Hex()
	return string
}