package services_test

import (
	"testing"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/stretchr/testify/assert"
)

var (
	authService *services.AuthServices
)

func init() {
	config, _ := nyatta_context.LoadConfig("..")
	logger, _ := services.NewLogger(config)
	authService = services.NewAuthService(logger, config)
}

func Test_AuthServices(t *testing.T) {
	var newUser *model.User
	var err error
	var jwt *string
	t.Run("should_create_new_fresh_jwt", func(t *testing.T) {
		newUser, err = userService.CreateUser(&model.NewUser{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "janedoe@email.com",
		})

		assert.Nil(t, err)

		jwt, err = authService.SignJWT(newUser)

		assert.Nil(t, err)
		assert.NotEqual(t, jwt, "")
	})

	t.Run("should_validate_jwt", func(t *testing.T) {
		token, err := authService.ValidateJWT(jwt)

		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.True(t, token.Valid)
	})
}
