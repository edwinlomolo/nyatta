package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
)

func Test_Amenity_service(t *testing.T) {
	user, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    "https://avatar.jpg",
	})
	property, err := propertyService.CreateProperty(&model.NewProperty{
		Name:       "Jonsaga Properties",
		Town:       "Upper Hill",
		PostalCode: "00500",
		CreatedBy:  user.ID,
	})

	assert.Nil(t, err)

	t.Run("should_add_property_amenity", func(t *testing.T) {
		amenity, err := amenityService.AddAmenity(&model.AmenityInput{
			Name:       "Home Fibre",
			Provider:   "Safaricom Home Internet Services",
			PropertyID: property.ID,
		})

		assert.Nil(t, err)
		assert.Equal(t, amenity.Name, "Home Fibre")
		assert.Equal(t, amenity.Provider, "Safaricom Home Internet Services")
	})

	t.Run("should_get_property_amenities", func(t *testing.T) {
		amenities, err := amenityService.PropertyAmenities(property.ID)

		assert.Nil(t, err)
		assert.Equal(t, len(amenities), 1)
	})

}
