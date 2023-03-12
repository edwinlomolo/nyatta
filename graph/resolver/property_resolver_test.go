package resolver

import (
	"fmt"
	"testing"
	"time"

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
			MinPrice   int
			MaxPrice   int
			PostalCode string
			ID         string
		}
	}

	var addPropertyUnit struct {
		AddPropertyUnit struct {
			Bathrooms int
			ID        string
		}
	}

	// authed server
	srv := makeAuthedGqlServer(true, ctx)

	// test user
	newUser, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
		Avatar:    "https://avatar.jpg",
	})
	if err != nil {
		t.Errorf("expected nil err got: %v", err)
	}

	t.Run("should_create_property", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { createProperty(input: { name: "%s", town: "%s", type: "%s", postalCode: "%s", minPrice: %v, maxPrice: %v, createdBy: "%s" }) { id, name, town, type, minPrice, maxPrice, postalCode } }`, "Oloolua Villas", "Ngong Hills", "Studio", "00208", 5000, 100000, newUser.ID)

		srv.MustPost(query, &createProperty)
		assert.Equal(t, createProperty.CreateProperty.Name, "Oloolua Villas")
		assert.Equal(t, createProperty.CreateProperty.PostalCode, "00208")
		assert.Equal(t, createProperty.CreateProperty.MinPrice, 5000)
		assert.Equal(t, createProperty.CreateProperty.MaxPrice, 100000)
		assert.Equal(t, createProperty.CreateProperty.Type, "Studio")
		assert.Equal(t, createProperty.CreateProperty.Town, "Ngong Hills")
	})

	t.Run("should_get_property_details", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name       string
				PostalCode string
				Type       string
				MinPrice   int
				MaxPrice   int
				Owner      struct {
					First_Name string
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, postalCode, minPrice, maxPrice, type, owner { first_name } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, createProperty.CreateProperty.Name, getProperty.GetProperty.Name)
		assert.Equal(t, createProperty.CreateProperty.PostalCode, getProperty.GetProperty.PostalCode)
		assert.Equal(t, createProperty.CreateProperty.MinPrice, getProperty.GetProperty.MinPrice)
		assert.Equal(t, createProperty.CreateProperty.MaxPrice, getProperty.GetProperty.MaxPrice)
		assert.Equal(t, createProperty.CreateProperty.Type, getProperty.GetProperty.Type)
		assert.Equal(t, getProperty.GetProperty.Owner.First_Name, "John")
	})

	t.Run("should_add_property_amenity", func(t *testing.T) {
		var amenity struct {
			AddAmenity struct {
				Name     string
				Provider string
			}
		}

		query := fmt.Sprintf(
			`mutation { addAmenity(input: {name: "Home Fibre", provider: "Safaricom Home Internet", propertyId: "%s"}) { name, provider } }`,
			createProperty.CreateProperty.ID,
		)
		srv.MustPost(query, &amenity)

	})

	t.Run("should_get_property_amenities", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name      string
				Amenities []struct {
					Name     string
					Provider string
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, amenities { name, provider } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, len(getProperty.GetProperty.Amenities), 1)
		assert.Equal(t, getProperty.GetProperty.Amenities[0].Name, "Home Fibre")
		assert.Equal(t, getProperty.GetProperty.Amenities[0].Provider, "Safaricom Home Internet")
	})

	t.Run("should_add_property_unit", func(t *testing.T) {
		query := fmt.Sprintf(
			`mutation { addPropertyUnit(input: {propertyId: %q, bathrooms: %v}) { id, bathrooms } }`,
			createProperty.CreateProperty.ID, 3,
		)

		srv.MustPost(query, &addPropertyUnit)

		assert.Equal(t, addPropertyUnit.AddPropertyUnit.Bathrooms, 3)
	})

	t.Run("should_query_property_unit(s)", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name  string
				Units []struct {
					Bathrooms int
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, units { bathrooms } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, len(getProperty.GetProperty.Units), 1)
		assert.Equal(t, getProperty.GetProperty.Units[0].Bathrooms, 3)
	})

	t.Run("should_add_unit_bedrooms", func(t *testing.T) {
		var addUnitBedrooms struct {
			AddUnitBedrooms []struct{ ID string }
		}
		query := fmt.Sprintf(
			`mutation { addUnitBedrooms(input: [{propertyUnitId: %q, bedroomNumber: %v, enSuite: %v, master: %v}]) { id } }`,
			addPropertyUnit.AddPropertyUnit.ID,
			1,
			true,
			true,
		)

		srv.MustPost(query, &addUnitBedrooms)

		assert.Equal(t, len(addUnitBedrooms.AddUnitBedrooms), 1)
	})

	t.Run("should_query_property_unit_bedrooms", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name  string
				Units []struct {
					Bedrooms []struct {
						BedroomNumber int
					}
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, units { bedrooms { bedroomNumber } } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, len(getProperty.GetProperty.Units[0].Bedrooms), 1)
		assert.Equal(t, getProperty.GetProperty.Units[0].Bedrooms[0].BedroomNumber, 1)
	})

	t.Run("should_add_property_unit_tenancy", func(t *testing.T) {
		when := time.Now().Format(time.RFC3339)
		var addPropertyUnitTenant struct {
			AddPropertyUnitTenant struct {
				StartDate string
			}
		}

		query := fmt.Sprintf(
			`mutation { addPropertyUnitTenant(input: {startDate: %q, endDate: %q, propertyUnitId: %q}) { startDate } }`,
			when,
			when,
			addPropertyUnit.AddPropertyUnit.ID,
		)

		srv.MustPost(query, &addPropertyUnitTenant)

		assert.NotEmpty(t, addPropertyUnitTenant.AddPropertyUnitTenant.StartDate)
	})

	t.Run("should_query_property_unit_tenancy", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name  string
				Units []struct {
					Tenancy []struct {
						StartDate string
					}
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, units { tenancy { startDate } } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, len(getProperty.GetProperty.Units[0].Tenancy), 1)
		assert.NotEmpty(t, getProperty.GetProperty.Units[0].Tenancy[0].StartDate)
	})

	t.Run("should_get_property_listings", func(t *testing.T) {
		var getListings struct {
			GetListings []struct {
				ID string
			}
		}
		query := fmt.Sprintf(
			`query { getListings(input: {town: %q, propertyType: %q, minPrice: %d, maxPrice: %d}) { id } }`,
			"Ngong Hills",
			"Studio",
			0,
			0,
		)

		srv.MustPost(query, &getListings)

		assert.Equal(t, len(getListings.GetListings), 0)
	})
}
