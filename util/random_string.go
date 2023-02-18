package util

import (
	"fmt"

	"github.com/rs/xid"
)

// randString - generate random string ids
func randString() string {
	guid := xid.New().String()
	return guid
}

// GenerateRandomEmail - generate random email address
func GenerateRandomEmail() string {
	randPrefix := randString()
	return fmt.Sprintf("%s@email.com", randPrefix)
}
