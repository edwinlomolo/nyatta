package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"
)

var (
	authService *AuthServices
)

func init() {
	config, _ := config.LoadConfig("..")
	logger, _ := NewLogger(config)
	authService = NewAuthService(logger, config)
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
	})
}
