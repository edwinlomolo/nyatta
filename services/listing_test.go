package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	town = "Ngong Hills"
)

func Test_Listing_Services(t *testing.T) {
	propertyService := NewPropertyService(queries, log.New())
	user, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    "https://avatar.jpg",
	})

	assert.Nil(t, err)

	var property *model.Property
	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, listingService.ServiceName(), "ListingServices")
	})

	t.Run("should_get_listings_with_all_correct_parameters", func(t *testing.T) {
		newProperty := []*model.NewProperty{
			{
				Name:       "Jonsaga Properties",
				Town:       town,
				PostalCode: "00500",
				Type:       "Studio",
				MinPrice:   5000,
				MaxPrice:   100000,
				CreatedBy:  user.ID,
			},
			{
				Name:       "Jonsaga Properties",
				Town:       town,
				PostalCode: "00500",
				Type:       "Bungalow",
				MinPrice:   5000,
				MaxPrice:   100000,
				CreatedBy:  user.ID,
			},
		}

		var err error
		for i := 0; i < len(newProperty); i++ {
			property, err = propertyService.CreateProperty(newProperty[i])
		}

		assert.Nil(t, err)
		assert.Equal(t, property.Name, "Jonsaga Properties")

		minPrice := 0
		maxPrice := 1000000
		propertyType := "Studio"
		listings, err := listingService.GetListings(model.ListingsInput{
			Town:         town,
			PropertyType: &propertyType,
			MinPrice:     &minPrice,
			MaxPrice:     &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 1)
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
		assert.Equal(t, len(listings), 1)
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
