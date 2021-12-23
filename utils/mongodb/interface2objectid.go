package mongodb

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Interface2ObjectId(v interface{}) primitive.ObjectID {
	data, err := bson.Marshal(v)
	if err != nil {
		log.Println("strcut to bson error ", err)
		return primitive.NilObjectID
	}
	var id primitive.ObjectID
	_ = bson.Unmarshal(data, &id)
	return id
}
