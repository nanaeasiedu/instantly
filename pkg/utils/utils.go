package utils

import (
	"math/rand"
	"time"
)

var allLeters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = allLeters[rand.Intn(len(allLeters))]
	}
	return string(b)
}
