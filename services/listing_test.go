package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"
)

func Test_Listing_Services(t *testing.T) {
	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, listingService.ServiceName(), "Listing")
	})

	t.Run("should_get_property_listings", func(t *testing.T) {
		minPrice := 0
		maxPrice := 0
		listings, err := listingService.GetListings(model.ListingsInput{
			Town:         "Ngong Hills",
			PropertyType: "studio",
			MinPrice:     &minPrice,
			MaxPrice:     &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 0)
	})
}
