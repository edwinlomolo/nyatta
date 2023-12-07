package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_property_service(t *testing.T) {
	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, propertyService.ServiceName(), "propertyClient")
	})

	t.Run("should_create_property", func(t *testing.T) {
	})

	t.Run("should_get_existing_property", func(t *testing.T) {
	})

	t.Run("should_error_finding_nonexistent_property", func(t *testing.T) {
	})

	t.Run("should_find_properties_by_town", func(t *testing.T) {
	})

	t.Run("should_find_properties_by_postal_code", func(t *testing.T) {
	})

	t.Run("should_get_properties_belonging_to_a_user", func(t *testing.T) {
	})

	t.Run("should_find_property_units", func(t *testing.T) {
	})
}
