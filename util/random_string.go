package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// randString - generate random string ids
func randString() string {
	uuid, _ := uuid.NewRandom()
	return uuid.String()
}

// GenerateRandomEmail - generate random email address
func GenerateRandomEmail() string {
	randPrefix := strings.Split(randString(), "-")[4]
	return fmt.Sprintf("%s@nyatta.app", randPrefix)
}
