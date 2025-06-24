package utils

import "math/rand"

var charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// GenerateSlug creates a random slug using the character set
func GenerateSlug(length uint8) string {
	slug := make([]rune, length)

	// Map bytes to character set
	for i := range slug {
		slug[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(slug)
}
