package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AuthServices(t *testing.T) {
	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, authService.ServiceName(), "AuthServices")
	})
}
