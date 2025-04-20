package utils

import (
	"fmt"
	"math/rand"
)

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func JoinCacheKey(key string, args ...interface{}) string {
	if len(args) == 0 {
		return key
	}
	for _, arg := range args {
		key += "_" + fmt.Sprintf("%v", arg)
	}
	key += ":"
	return key
}
