package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"
)

func Test_property_service(t *testing.T) {
	propertyService := NewPropertyService(store, logger)

	var property *model.Property

	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, propertyService.ServiceName(), "PropertyServices")
	})

	t.Run("should_create_property", func(t *testing.T) {
		newProperty := &model.NewProperty{
			Name:       "Jonsaga Properties",
			Town:       "Upper Hill",
			PostalCode: "00500",
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
}
