package utils

import (
	"log"
	"strconv"
)

func String2Uint64(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Printf("string to uint64 fail, err=[%v]", err)
	}
	return res
}

func Uint642String(i uint64) string {
	res := strconv.FormatUint(i, 10)
	return res
}
