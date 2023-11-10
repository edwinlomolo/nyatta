package util

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_string_randomness(t *testing.T) {
	t.Run("should_generate_random_email", func(t *testing.T) {

		email := GenerateRandomEmail()
		match, _ := regexp.MatchString("([a-zA-Z0-9]+)@nyatta.app", email)
		assert.True(t, match)
	})
}
