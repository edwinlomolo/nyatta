package resolver

import (
	"testing"
)

func Test_Property_Resolver(t *testing.T) {
	// test user
	_, err := userService.FindUserByPhone("+254712345678")
	if err != nil {
		t.Errorf("expected nil err got: %v", err)
	}

	t.Run("should_create_property", func(t *testing.T) {
	})

	t.Run("should_get_property_details", func(t *testing.T) {
	})

	t.Run("should_add_property_unit", func(t *testing.T) {
	})

	t.Run("should_query_property_unit(s)", func(t *testing.T) {
	})

	t.Run("should_add_unit_bedrooms", func(t *testing.T) {
	})

	t.Run("should_query_property_unit_bedrooms", func(t *testing.T) {
	})

	t.Run("should_add_property_unit_tenancy", func(t *testing.T) {
	})

	t.Run("should_query_property_unit_tenancy", func(t *testing.T) {
	})
}
