package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Time2Timestamp(t time.Time) primitive.Timestamp {
	return primitive.Timestamp{T: uint32(t.Unix())}
}

func Timestamp2Time(ts primitive.Timestamp) time.Time {
	return time.Unix(int64(ts.T), 0)
}
