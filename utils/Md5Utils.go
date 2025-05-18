package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(source string) string {
	return Md5WithSalt(source, "")
}

func Md5WithSalt(source string, salt string) string {
	return Md5WithSaltAndIterations(source, salt, 1)
}

func Md5WithSaltAndIterations(source string, salt string, iterations int) string {
	if iterations < 1 {
		iterations = 1
	}
	hash := doMd5(source + salt)
	for i := 0; i < iterations; i++ {
		hash = doMd5(hash)
	}
	return hash
}

func doMd5(source string) string {
	data := []byte(source)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
