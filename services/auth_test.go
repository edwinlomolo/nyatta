package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
)

func Test_AuthServices(t *testing.T) {
	var newUser *model.User
	var err error
	var jwt *string

	t.Run("should_get_service_name", func(t *testing.T) {
		assert.Equal(t, authService.ServiceName(), "AuthServices")
	})

	t.Run("should_create_new_fresh_jwt", func(t *testing.T) {
		newUser, err = userService.CreateUser(&model.NewUser{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     util.GenerateRandomEmail(),
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
