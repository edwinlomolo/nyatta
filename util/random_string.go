package util

import (
	"fmt"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randString - generate random string ids
func randString() string {
	s := make([]rune, 5)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// GenerateRandomEmail - generate random email address
func GenerateRandomEmail() string {
	return fmt.Sprintf("%s@email.com", randString())
}
