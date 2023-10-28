package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"
)

func Test_Amenity_service(t *testing.T) {
	user, err := userService.FindUserByPhone("+2549869240")
	property, err := propertyService.CreateProperty(&model.NewProperty{
		Name:       "Jonsaga Properties",
		Town:       "Upper Hill",
		PostalCode: "00500",
		CreatedBy:  user.ID,
	})

	assert.Nil(t, err)
	assert.Equal(t, property.Name, "Jonsaga Properties")

	t.Run("should_add_unit_amenity", func(t *testing.T) {
	})

	t.Run("should_get_unit_amenities", func(t *testing.T) {
	})

}
