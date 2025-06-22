package utils

import (
	"math/rand"
	"time"
)

var charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.NewSource(time.Now().Unix())
}

func GenerateSlug(length uint8) string {
	slug := make([]rune, length)
	for i := range slug {
		slug[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(slug)
}
