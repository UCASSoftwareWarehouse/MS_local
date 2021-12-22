package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func Interface2ObjectId(v interface{}) primitive.ObjectID {
	data, err := bson.Marshal(v)
	if err != nil {
		log.Fatal("strcut to bson error ", err)
	}
	var id primitive.ObjectID
	_ = bson.Unmarshal(data, &id)
	return id
}
