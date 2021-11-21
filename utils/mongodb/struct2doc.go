package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

//https://studygolang.com/articles/2552
//https://gist.github.com/divjotarora/06c5188138456070cee26024f223b3ee
//https://pkg.go.dev/github.com/mongodb/mongo-go-driver/bson#Marshal
func Struct2Bson(v interface{}) bson.D {
	//fmt.Println(v)
	data, err := bson.Marshal(v)
	//fmt.Println("%q", data)
	if err != nil {
		log.Fatal("strcut to bson error", err)
	}
	var doc bson.D
	_ = bson.Unmarshal(data, &doc)
	//fmt.Println("doc:", doc)
	return doc
}
