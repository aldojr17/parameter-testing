package util

import (
	"os"
	"time"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GenerateCurrentTimestamp() int64 {
	return time.Now().Unix()
}
