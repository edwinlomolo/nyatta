package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
)

func Test_property_service(t *testing.T) {
	propertyService := NewPropertyService(store, logger)
	user, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
	})
	assert.Nil(t, err)

	var property *model.Property

	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, propertyService.ServiceName(), "PropertyServices")
	})

	t.Run("should_create_property", func(t *testing.T) {
		newProperty := &model.NewProperty{
			Name:       "Jonsaga Properties",
			Town:       "Upper Hill",
			PostalCode: "00500",
			CreatedBy:  user.ID,
		}
		var err error

		property, err = propertyService.CreateProperty(newProperty)

		assert.Nil(t, err)
		assert.Equal(t, property.Name, "Jonsaga Properties")
		assert.Equal(t, property.Town, "Upper Hill")
		assert.Equal(t, property.PostalCode, "00500")
		assert.NotEmpty(t, property.ID)
	})

	t.Run("should_get_existing_property", func(t *testing.T) {
		foundProperty, err := propertyService.GetProperty(property.ID)

		assert.Nil(t, err)
		assert.Equal(t, foundProperty.ID, property.ID)
	})

	t.Run("should_error_finding_nonexistent_property", func(t *testing.T) {
		foundProperty, err := propertyService.GetProperty("erkhlshf")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Property does not exist")
		assert.Nil(t, foundProperty)
	})

	t.Run("should_find_properties_by_town", func(t *testing.T) {
		foundProperties, err := propertyService.FindByTown(property.Town)

		assert.Nil(t, err)
		assert.Equal(t, len(foundProperties), 1)
	})

	t.Run("should_find_properties_by_postal_code", func(t *testing.T) {
		foundProperties, err := propertyService.FindByPostalCode(property.PostalCode)

		assert.Nil(t, err)
		assert.Equal(t, len(foundProperties), 1)
	})

	t.Run("should_add_property_amenity", func(t *testing.T) {
		amenity, err := propertyService.AddAmenity(&model.AmenityInput{
			Name:       "Home Fibre",
			Provider:   "Safaricom Home Internet Services",
			PropertyID: property.ID,
		})

		assert.Nil(t, err)
		assert.Equal(t, amenity.Name, "Home Fibre")
		assert.Equal(t, amenity.Provider, "Safaricom Home Internet Services")

		// get property amenities
		property, err := propertyService.GetProperty(property.ID)

		assert.Nil(t, err)
		assert.Equal(t, len(property.Amenities), 1)
	})

	t.Run("should_get_properties_belonging_to_a_user", func(t *testing.T) {
		user, err := userService.FindById(user.ID)
		assert.Nil(t, err)
		assert.Equal(t, len(user.Properties), 1)
	})
}
