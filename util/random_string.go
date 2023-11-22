package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// randString - generate random string ids
func randString() string {
	guid, _ := uuid.NewRandom()
	return guid.String()
}

// GenerateRandomEmail - generate random email address
func GenerateRandomEmail() string {
	randPrefix := strings.ReplaceAll(randString(), "-", "")
	return fmt.Sprintf("%s@nyatta.app", randPrefix)
}
