package util_test

import (
	"regexp"
	"testing"

	"github.com/3dw1nM0535/nyatta/util"
)

func Test_string_randomness(t *testing.T) {
	t.Run("should_generate_random_email", func(t *testing.T) {

		email := util.GenerateRandomEmail()
		match, _ := regexp.MatchString("([a-zA-Z]+)@email.com", email)
		if !match {
			t.Errorf("expected %s to contain email example", email)
		}
	})
}
