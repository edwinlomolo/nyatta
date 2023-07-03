package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Listing_Services(t *testing.T) {
	//	propertyService := NewPropertyService(queries, log.New())
	//	user, _ := userService.CreateUser(&model.NewUser{
	//		FirstName: "John",
	//		LastName:  "Doe",
	//		Email:     util.GenerateRandomEmail(),
	//		Avatar:    "https://avatar.jpg",
	//		Phone:     "+25427348342",
	//	})
	//
	//	// Seed some property(s)
	//	newProperty := []*model.NewProperty{
	//		{
	//			Name:       "Kadong Villa",
	//			Town:       "Karen",
	//			PostalCode: "10345",
	//			Type:       "Home",
	//			CreatedBy:  user.ID,
	//		},
	//		{
	//			Name:       "Jonsaga Properties",
	//			Town:       "Ngong Hills",
	//			PostalCode: "00500",
	//			Type:       "Flat",
	//			CreatedBy:  user.ID,
	//		},
	//	}
	//
	//	for i := 0; i < len(newProperty); i++ {
	//		_, _ = propertyService.CreateProperty(newProperty[i])
	//	}

	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, listingService.ServiceName(), "ListingServices")
	})

	t.Run("should_get_vacant_listings", func(t *testing.T) {
	})
}
