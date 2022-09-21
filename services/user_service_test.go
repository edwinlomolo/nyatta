package services

import (
	"testing"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	userService   *UserServices
	store         *gorm.DB
	configuration *nyatta_context.Config
)

func init() {
	configuration, _ = nyatta_context.LoadConfig("..")
	logger, _ = NewLogger(configuration)
	store, _ = nyatta_context.OpenDB(configuration, logger)
	userService = NewUserService(store, logger, configuration)
}

func Test_User_Services(t *testing.T) {
	var newUser *model.User
	var err error

	t.Run("should_create_new_user", func(t *testing.T) {
		newUser, err = userService.CreateUser(&model.NewUser{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@email.com",
		})

		assert.Nil(t, err)
		assert.Equal(t, newUser.FirstName, "John")
		assert.Equal(t, newUser.LastName, "Doe")
	})

	t.Run("should_sign_in_user", func(t *testing.T) {
		token, err := userService.SignIn(&model.NewUser{FirstName: "John", LastName: "Doe", Email: util.GenerateRandomEmail()})

		assert.Nil(t, err)
		assert.NotEqual(t, token, "")

		tokenValid, err := userService.ValidateToken(token)

		assert.Nil(t, err)
		assert.True(t, tokenValid.Valid)
	})

	t.Run("should_get_existing_user_by_id", func(t *testing.T) {
		foundUser, err := userService.FindById(newUser.ID)

		assert.Equal(t, foundUser.FirstName, "John")
		assert.Equal(t, foundUser.LastName, "Doe")
		assert.Nil(t, err)
	})

	t.Run("should_get_existing_user_by_email", func(t *testing.T) {
		foundUser, err := userService.FindByEmail(newUser.Email)

		assert.Nil(t, err)
		assert.Equal(t, newUser.Email, foundUser.Email)
	})

	t.Run("should_get_service_name_called", func(t *testing.T) {
		assert.Equal(t, userService.ServiceName(), "UserServices")
	})
}
