package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
)

func TestObjectId2String(t *testing.T) {
	temp_id := primitive.NewObjectID()
	log.Println(temp_id)
	log.Println(len(temp_id))
	string_id := ObjectId2String(temp_id)
	log.Println(string_id)
	log.Println(len(string_id))
	log.Println(String2ObjectId(string_id))
}
