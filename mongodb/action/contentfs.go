package action

import (
	"MS_Local/mongodb"
	mongodb2 "MS_Local/utils/mongodb"
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UploadGridFile(file, filename string) (*primitive.ObjectID, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("open file error, err=[%v]", err)
		return nil, err
	}
	bucket, err := gridfs.NewBucket(
		mongodb.GridFS,
	)
	if err != nil {
		log.Printf("create bucket error, err=[%v]", err)
		return nil, err
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	if err != nil {
		log.Printf("create upload stream error, err=[%v]", err)
		return nil, err
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Printf("write file to gridfs error, err=[%v]", err)
		return nil, err
	}
	log.Printf("Write file %s to GRIDFS was successful. File size: %d b\n", filename, fileSize)
	fid := uploadStream.FileID.(primitive.ObjectID)
	return &fid, nil
}

func DownloadGridFile(fname string, id string) error {
	oid, err := mongodb2.String2ObjectId(id)
	if err != nil {
		return err
	}
	bucket, _ := gridfs.NewBucket(
		mongodb.GridFS,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStream(oid, &buf)
	if err != nil {
		log.Printf("download to stream fail, err=[%v]", err)
		return err
	}
	ioutil.WriteFile(fname, buf.Bytes(), 0644)
	log.Printf("Download file %s from gridfs success, File size is: %v \n", fname, dStream)
	return nil

}

func DeleteGridFile(id string) error {
	fsFiles := mongodb.GridFS.Collection("fs.files")
	tmp_id, err := mongodb2.String2ObjectId(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": tmp_id}
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	deleteResult, err := fsFiles.DeleteOne(context.Background(), filter, opts)
	if err != nil {
		log.Printf("delete gridfile error, err=[%v]", err)
		return err
	}
	log.Printf("Delete GridFile: %v", deleteResult)
	return nil
}
