package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
)

func Test_Unit_Services(t *testing.T) {
	var unit *model.PropertyUnit
	user, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
	})
	assert.Nil(t, err)

	property, err := propertyService.CreateProperty(&model.NewProperty{
		Name:       "Jonsaga Properties",
		Town:       "Upper Hill",
		PostalCode: "00500",
		CreatedBy:  user.ID,
	})
	assert.Nil(t, err)

	t.Run("should_add_property_unit", func(t *testing.T) {
		newUnit := &model.PropertyUnitInput{
			PropertyID: property.ID,
			Bathrooms:  3,
		}

		var err error
		unit, err = unitService.AddPropertyUnit(newUnit)

		assert.Nil(t, err)
		assert.Equal(t, unit.Bathrooms, 3)
	})

	t.Run("should_add_property_unit_bedrooms", func(t *testing.T) {
		bedroom1 := &model.UnitBedroomInput{
			BedroomNumber:  1,
			PropertyUnitID: unit.ID,
			EnSuite:        true,
			Master:         true,
		}
		bedroom2 := &model.UnitBedroomInput{
			BedroomNumber:  2,
			PropertyUnitID: unit.ID,
			EnSuite:        false,
			Master:         true,
		}
		newBedroom := []*model.UnitBedroomInput{
			bedroom1,
			bedroom2,
		}

		insertedBedroom, err := unitService.AddUnitBedrooms(newBedroom)

		assert.Nil(t, err)
		assert.Equal(t, len(insertedBedroom), 2)
	})

	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, unitService.ServiceName(), "UnitServices")
	})
}
