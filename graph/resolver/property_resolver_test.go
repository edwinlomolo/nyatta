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
			PostalCode string
			ID         string
		}
	}

	// authed server
	srv := makeAuthedGqlServer(true, ctx)

	// test user
	newUser, err := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     util.GenerateRandomEmail(),
	})
	if err != nil {
		t.Errorf("expected nil err got: %v", err)
	}

	t.Run("should_create_property", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { createProperty(input: { name: "%s", town: "%s", postalCode: "%s", createdBy: "%s" }) { id, name, town, postalCode } }`, "Oloolua Villas", "Ngong Hills", "00208", newUser.ID)

		srv.MustPost(query, &createProperty)
		assert.Equal(t, createProperty.CreateProperty.Name, "Oloolua Villas")
		assert.Equal(t, createProperty.CreateProperty.PostalCode, "00208")
		assert.Equal(t, createProperty.CreateProperty.Town, "Ngong Hills")
	})

	t.Run("should_get_property_details", func(t *testing.T) {
		var getProperty struct {
			GetProperty struct {
				Name       string
				PostalCode string
				Owner      struct {
					First_Name string
				}
			}
		}

		query := `query ($id: ID!) { getProperty(id: $id) { name, postalCode, owner { first_name } } }`

		srv.MustPost(query, &getProperty, client.Var("id", createProperty.CreateProperty.ID))

		assert.Equal(t, createProperty.CreateProperty.Name, getProperty.GetProperty.Name)
		assert.Equal(t, createProperty.CreateProperty.PostalCode, getProperty.GetProperty.PostalCode)
		assert.Equal(t, getProperty.GetProperty.Owner.First_Name, "John")
	})
}
