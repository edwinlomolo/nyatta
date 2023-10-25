package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
)

func Test_Amenity_service(t *testing.T) {
	avatar := "https://avatar.jpg"
	user, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    &avatar,
	})
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
