package mongodb

import (
	"MS_Local/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)


func InitMongo() (*mongo.Database, error) {
	var (
		client   *mongo.Client
		mongoURL = config.MongoUrl
	)

	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))

	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("connect error", err)
	}
	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("ping error", err)
	}
	db := client.Database("Files")
	return db, err
}
func GetCollectionFromMongo(DB *mongo.Database, collectionName string) *mongo.Collection {
	// 获取数据库和集合
	collection := DB.Collection(collectionName)
	return collection
}
