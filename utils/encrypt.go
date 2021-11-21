package utils

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)


// Encrypt 密码加密算法
func Encrypt(data string) string {
	if data == "" {
		log.Fatal("password is empty")
	}
	h := md5.New()
	// 加yan操作
	dataBytes := []byte(data)
	if dataBytes[0] > 5 {
		dataBytes[0] -= 4
	} else {
		dataBytes[0] += 3
	}

	h.Write(dataBytes)
	return hex.EncodeToString(h.Sum(nil))
}

