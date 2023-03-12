package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"
)

const (
	town = "Ngong Hills"
)

func Test_Listing_Services(t *testing.T) {
	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, listingService.ServiceName(), "ListingServices")
	})

	t.Run("should_get_listings_with_all_correct_parameters", func(t *testing.T) {
		minPrice := 0
		maxPrice := 1000
		propertyType := "Studio"
		listings, err := listingService.GetListings(model.ListingsInput{
			Town:         town,
			PropertyType: &propertyType,
			MinPrice:     &minPrice,
			MaxPrice:     &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 0)
	})

	t.Run("should_get_listings_with_zero_pricing", func(t *testing.T) {
		minPrice := 0
		maxPrice := 0
		propertyType := "Bungalow"
		listings, err := listingService.GetListings(model.ListingsInput{
			Town:         town,
			PropertyType: &propertyType,
			MinPrice:     &minPrice,
			MaxPrice:     &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 0)
	})

	t.Run("should_get_listings_without_property_type_param", func(t *testing.T) {
		minPrice := 0
		maxPrice := 10000
		propertyType := ""

		listings, err := listingService.GetListings(model.ListingsInput{
			Town:         town,
			MinPrice:     &minPrice,
			MaxPrice:     &maxPrice,
			PropertyType: &propertyType,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 0)
	})
}
