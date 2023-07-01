package resolver

import (
	"fmt"
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func Test_Resolver_User(t *testing.T) {
	var signIn struct {
		SignIn struct {
			Token string
		}
	}
	var user *model.User

	// authed server
	srv := makeAuthedGqlServer(false, ctx)

	t.Run("resolver_should_sign_in_user", func(t *testing.T) {

		query := fmt.Sprintf(`mutation { signIn (input: { first_name: %q, last_name: %q, email: %q, avatar: %q, phone: %q }) { token } }`, "Jane", "Doe", util.GenerateRandomEmail(), "https://avatar.jpg", "+254002397075")

		srv.MustPost(query, &signIn)

		assert.NotEqual(t, signIn.SignIn.Token, "")
	})
	t.Run("resolver_should_get_user", func(t *testing.T) {
		var getUser struct {
			GetUser struct {
				Email      string
				First_Name string
			}
		}
		var err error

		email := util.GenerateRandomEmail()
		user, err = userService.CreateUser(&model.NewUser{FirstName: "John", LastName: "Doe", Email: email, Avatar: "https://avatar.jpg", Phone: "+254829639846"})
		if err != nil {
			t.Errorf("expected nil err got %v", err)
		}
		srv.MustPost(`query ($email: String!) { getUser (email: $email) { email, first_name  } }`, &getUser, client.Var("email", email))

		assert.Equal(t, getUser.GetUser.Email, email)
		assert.Equal(t, getUser.GetUser.First_Name, "John")
	})

	t.Run("resolver_should_get_properties_belonging_to_user", func(t *testing.T) {
		var getUser struct {
			GetUser struct {
				Properties []struct{ Name string }
			}
		}

		_, err = propertyService.CreateProperty(&model.NewProperty{
			Name:       "Ngong Hills Agency",
			PostalCode: "00208",
			Town:       "Ngong Hills",
			CreatedBy:  user.ID,
		})
		assert.Nil(t, err)

		srv.MustPost(`query ($email: String!) { getUser (email: $email) { properties { name } } }`, &getUser, client.Var("email", user.Email))

		assert.Equal(t, len(getUser.GetUser.Properties), 1)
	})
}
