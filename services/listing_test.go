package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Listing_Services(t *testing.T) {
	propertyService := NewPropertyService(queries, log.New())
	user, _ := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    "https://avatar.jpg",
		Phone:     "+25427348342",
	})

	// Seed some property(s)
	newProperty := []*model.NewProperty{
		{
			Name:       "Kadong Villa",
			Town:       "Karen",
			PostalCode: "10345",
			Type:       "Home",
			MinPrice:   40000,
			MaxPrice:   0,
			CreatedBy:  user.ID,
		},
		{
			Name:       "Jonsaga Properties",
			Town:       "Ngong Hills",
			PostalCode: "00500",
			Type:       "Flat",
			MinPrice:   5000,
			MaxPrice:   35000,
			CreatedBy:  user.ID,
		},
	}

	for i := 0; i < len(newProperty); i++ {
		_, _ = propertyService.CreateProperty(newProperty[i])
	}

	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, listingService.ServiceName(), "ListingServices")
	})

	t.Run("should_get_listings_with_all_correct_parameters", func(t *testing.T) {
		var listings []*model.Property
		var err error
		var minPrice int = 40000
		var maxPrice int = 0

		listings, err = listingService.GetListings(model.ListingsInput{
			Town:     "Karen",
			MinPrice: &minPrice,
			MaxPrice: &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 1)
		assert.Equal(t, listings[0].Town, "Karen")

		minPrice = 5000
		maxPrice = 35000
		listings, _ = listingService.GetListings(model.ListingsInput{
			Town:     "Ngong Hills",
			MinPrice: &minPrice,
			MaxPrice: &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 1)
		assert.Equal(t, listings[0].Town, "Ngong Hills")
	})

	t.Run("should_get_listings_with_zero_pricing", func(t *testing.T) {
		minPrice := 0
		maxPrice := 0
		listings, err := listingService.GetListings(model.ListingsInput{
			Town:     "Ngong Hills",
			MinPrice: &minPrice,
			MaxPrice: &maxPrice,
		})

		assert.Nil(t, err)
		assert.Equal(t, len(listings), 1)
	})
}
