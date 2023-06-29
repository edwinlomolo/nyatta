package services

import (
	"fmt"
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_property_service(t *testing.T) {
	propertyService := NewPropertyService(queries, log.New())
	user, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    "https://avatar.jpg",
		Phone:     "+254828992034",
	})

	var property *model.Property

	assert.Nil(t, err)

	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, propertyService.ServiceName(), "PropertyServices")
	})

	t.Run("should_create_property", func(t *testing.T) {
		newProperty := &model.NewProperty{
			Name:       "Jonsaga Properties",
			Town:       "Upper Hill",
			PostalCode: "00500",
			Type:       "Apartment",
			CreatedBy:  user.ID,
		}
		var err error

		property, err = propertyService.CreateProperty(newProperty)

		assert.Nil(t, err)
		assert.Equal(t, property.Name, "Jonsaga Properties")
		assert.Equal(t, property.Town, "Upper Hill")
		assert.Equal(t, property.PostalCode, "00500")
		assert.Equal(t, property.Type, "Apartment")
		assert.NotEmpty(t, property.ID)
	})

	t.Run("should_get_existing_property", func(t *testing.T) {
		foundProperty, err := propertyService.GetProperty(property.ID)

		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprint(foundProperty.ID), property.ID)
	})

	t.Run("should_error_finding_nonexistent_property", func(t *testing.T) {
		foundProperty, err := propertyService.GetProperty("97304702")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Property does not exist")
		assert.Nil(t, foundProperty)
	})

	t.Run("should_find_properties_by_town", func(t *testing.T) {
		foundProperties, err := propertyService.FindByTown(property.Town)

		assert.Nil(t, err)
		assert.Equal(t, len(foundProperties), 0)
	})

	t.Run("should_find_properties_by_postal_code", func(t *testing.T) {
		foundProperties, err := propertyService.FindByPostalCode(property.PostalCode)

		assert.Nil(t, err)
		assert.Equal(t, len(foundProperties), 0)
	})

	t.Run("should_get_properties_belonging_to_a_user", func(t *testing.T) {
		userProperties, err := propertyService.PropertiesCreatedBy(user.ID)

		assert.Nil(t, err)
		assert.Equal(t, len(userProperties), 1)
	})

	t.Run("should_find_property_units", func(t *testing.T) {
	})
}
