package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Listing_Services(t *testing.T) {
	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, listingService.ServiceName(), "listingClient")
	})

	t.Run("should_get_vacant_listings", func(t *testing.T) {
	})
}
