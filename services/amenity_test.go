package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Amenity_service(t *testing.T) {
	t.Run("should_add_unit_amenity", func(t *testing.T) {
	})

	t.Run("should_get_unit_amenities", func(t *testing.T) {
	})

	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, amenityService.ServiceName(), "amenityClient")
	})

}
