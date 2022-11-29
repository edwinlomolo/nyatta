package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Unit_Services(t *testing.T) {
	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, unitService.ServiceName(), "UnitServices")
	})
}
