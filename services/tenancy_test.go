package services

import (
	"testing"
	"time"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
)

func Test_Tenancy_Services(t *testing.T) {
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

	unit, err := unitService.AddPropertyUnit(&model.PropertyUnitInput{
		PropertyID: property.ID,
		Bathrooms:  3,
	})
	assert.Nil(t, err)

	t.Run("should_add_unit_tenant", func(t *testing.T) {

		newTenant := &model.TenancyInput{
			StartDate:      time.Now(),
			EndDate:        &time.Time{},
			PropertyUnitID: unit.ID,
		}

		insertedTenant, err := tenancyService.AddUnitTenancy(newTenant)

		assert.Nil(t, err)
		assert.NotEmpty(t, insertedTenant.StartDate)
		assert.NotEmpty(t, insertedTenant.EndDate)
	})

	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, tenancyService.ServiceName(), "TenancyServices")
	})
}
