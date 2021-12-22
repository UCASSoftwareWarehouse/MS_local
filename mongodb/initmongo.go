package mongodb

import (
	"MS_Local/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Softwaredb *mongo.Database
var BinaryCol *mongo.Collection
var CodeCol *mongo.Collection

func InitMongo() error {
	log.Printf("init mongo...")
	var (
		client   *mongo.Client
		mongoURL = fmt.Sprintf("%s", config.Conf.MongodbAddr)
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
	Softwaredb = db
	initCollection(db)
	return err
}

func initCollection(DB *mongo.Database) {
	cnames, err := DB.ListCollectionNames(context.Background(),bson.D{})
	if(err!=nil){
		log.Printf("list collections error, err=[%v]", err)
	}
	var flags = [2]int{0,0}
	for _, name := range(cnames){
		if name==config.Conf.MongodbBinary{
			flags[0]=1
		}
		if name==config.Conf.MongodbCode{
			flags[1]=1
		}
	}
	//crete collection
	if flags[0]==0{
		err:=DB.CreateCollection(context.Background(), config.Conf.MongodbBinary)
		if(err!=nil){
			log.Printf("create collection err, err=[%v]", err)
		}
	}
	if flags[1]==0{
		err:= DB.CreateCollection(context.Background(), config.Conf.MongodbCode)
		if(err!=nil){
			log.Printf("create collection err, err=[%v]", err)
		}
	}
	bc := DB.Collection(config.Conf.MongodbBinary)
	BinaryCol = bc
	cc := DB.Collection(config.Conf.MongodbCode)
	CodeCol = cc
}

func GetCollectionFromMongo(DB *mongo.Database, collectionName string) *mongo.Collection {
	// 获取数据库和集合
	collection := DB.Collection(collectionName)
	return collection
}
