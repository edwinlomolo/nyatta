package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_property_service(t *testing.T) {
	propertyService := NewPropertyService(store, logger)

	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, propertyService.ServiceName(), "Property Service")
	})
}
