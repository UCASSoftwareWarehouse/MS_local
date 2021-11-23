package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func Time2Timestamp(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   0,
	}
}
