package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func String2ObjectId(s string) primitive.ObjectID {
	objId, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Fatal(err)
	}
	return objId
}

func ObjectId2String(id primitive.ObjectID) string {
	string := id.Hex()
	return string
}