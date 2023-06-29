package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Unit_Services(t *testing.T) {
	t.Run("should_add_property_unit", func(t *testing.T) {
	})

	t.Run("should_add_property_unit_bedrooms", func(t *testing.T) {
	})

	t.Run("should_get_unit_bedrooms", func(t *testing.T) {
	})

	t.Run("should_add_unit_tenant", func(t *testing.T) {
	})

	t.Run("should_get_unit_tenancy", func(t *testing.T) {
	})

	t.Run("should_return_service_name", func(t *testing.T) {
		assert.Equal(t, unitService.ServiceName(), "UnitServices")
	})
}
