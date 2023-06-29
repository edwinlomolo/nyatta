package resolver

import (
	"fmt"
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/assert"
)

func Test_Property_Resolver(t *testing.T) {
	var createProperty struct {
		CreateProperty struct {
			Name       string
			Town       string
			Type       string
			PostalCode string
			ID         string
		}
	}

	// authed server
	srv := makeAuthedGqlServer(false, ctx)

	// test user
	newUser, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    "https://avatar.jpg",
		Phone:     "+254712345678",
	})
	if err != nil {
		t.Errorf("expected nil err got: %v", err)
	}

	t.Run("should_create_property", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { createProperty(input: { name: "%s", town: "%s", type: "%s", postalCode: "%s", createdBy: "%s" }) { id, name, town, type, postalCode } }`, "Oloolua Villas", "Ngong Hills", "Studio", "00208", newUser.ID)

		srv.MustPost(query, &createProperty)
		assert.Equal(t, createProperty.CreateProperty.Name, "Oloolua Villas")
		assert.Equal(t, createProperty.CreateProperty.PostalCode, "00208")
		assert.Equal(t, createProperty.CreateProperty.Type, "Studio")
		assert.Equal(t, createProperty.CreateProperty.Town, "Ngong Hills")
	})

	t.Run("should_get_property_details", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name       string
				PostalCode string
				Type       string
				Owner      struct {
					First_Name string
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, postalCode, type, owner { first_name } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, createProperty.CreateProperty.Name, getProperty.GetProperty.Name)
		assert.Equal(t, createProperty.CreateProperty.PostalCode, getProperty.GetProperty.PostalCode)
		assert.Equal(t, createProperty.CreateProperty.Type, getProperty.GetProperty.Type)
		assert.Equal(t, getProperty.GetProperty.Owner.First_Name, "John")
	})

	t.Run("should_add_property_unit", func(t *testing.T) {
		//query := fmt.Sprintf(
		//	`mutation { addPropertyUnit(input: {propertyId: %q, bathrooms: %v}) { id, bathrooms } }`,
		//	createProperty.CreateProperty.ID, 3,
		//)

		//srv.MustPost(query, &addPropertyUnit)

		//assert.Equal(t, addPropertyUnit.AddPropertyUnit.Bathrooms, 3)
	})

	t.Run("should_query_property_unit(s)", func(t *testing.T) {
		//var getProperty struct {
		//	GetProperty struct {
		//		Name  string
		//		Units []struct {
		//			Bathrooms int
		//		}
		//	}
		//}

		//query := `query ($id: ID!) { getProperty(id: $id) { name, units { bathrooms } } }`

		//srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		//assert.Equal(t, len(getProperty.GetProperty.Units), 1)
		//assert.Equal(t, getProperty.GetProperty.Units[0].Bathrooms, 3)
	})

	t.Run("should_add_unit_bedrooms", func(t *testing.T) {
		//var addUnitBedrooms struct {
		//	AddUnitBedrooms []struct{ ID string }
		//}
		//query := fmt.Sprintf(
		//	`mutation { addUnitBedrooms(input: [{propertyUnitId: %q, bedroomNumber: %v, enSuite: %v, master: %v}]) { id } }`,
		//	addPropertyUnit.AddPropertyUnit.ID,
		//	1,
		//	true,
		//	true,
		//)

		//srv.MustPost(query, &addUnitBedrooms)

		//assert.Equal(t, len(addUnitBedrooms.AddUnitBedrooms), 1)
	})

	t.Run("should_query_property_unit_bedrooms", func(t *testing.T) {
		//var getProperty struct {
		//	GetProperty struct {
		//		Name  string
		//		Units []struct {
		//			Bedrooms []struct {
		//				BedroomNumber int
		//			}
		//		}
		//	}
		//}

		//query := `query ($id: ID!) { getProperty(id: $id) { name, units { bedrooms { bedroomNumber } } } }`

		//srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		//assert.Equal(t, len(getProperty.GetProperty.Units[0].Bedrooms), 1)
		//assert.Equal(t, getProperty.GetProperty.Units[0].Bedrooms[0].BedroomNumber, 1)
	})

	t.Run("should_add_property_unit_tenancy", func(t *testing.T) {
		///when := time.Now().Format(time.RFC3339)
		///var addPropertyUnitTenant struct {
		///	AddPropertyUnitTenant struct {
		///		StartDate string
		///	}
		///}

		///query := fmt.Sprintf(
		///	`mutation { addPropertyUnitTenant(input: {startDate: %q, endDate: %q, propertyUnitId: %q}) { startDate } }`,
		///	when,
		///	when,
		///	addPropertyUnit.AddPropertyUnit.ID,
		///)

		///srv.MustPost(query, &addPropertyUnitTenant)

		///assert.NotEmpty(t, addPropertyUnitTenant.AddPropertyUnitTenant.StartDate)
	})

	t.Run("should_query_property_unit_tenancy", func(t *testing.T) {
		//var getProperty struct {
		//	GetProperty struct {
		//		Name  string
		//		Units []struct {
		//			Tenancy []struct {
		//				StartDate string
		//			}
		//		}
		//	}
		//}

		//query := `query ($id: ID!) { getProperty(id: $id) { name, units { tenancy { startDate } } } }`

		//srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		//assert.Equal(t, len(getProperty.GetProperty.Units[0].Tenancy), 1)
		//assert.NotEmpty(t, getProperty.GetProperty.Units[0].Tenancy[0].StartDate)
	})
}
