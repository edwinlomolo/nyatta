package services_test

import (
	"testing"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	userService   *services.UserServices
	store         *gorm.DB
	logger        *zap.SugaredLogger
	configuration *nyatta_context.Config
)

func init() {
	configuration, _ = nyatta_context.LoadConfig("..")
	logger, _ = services.NewLogger(configuration)
	store, _ = nyatta_context.OpenDB(configuration, logger)
	userService = services.NewUserService(store, logger)
}

func Test_User_Services(t *testing.T) {
	var newUser *model.User
	t.Run("should_create_new_user", func(t *testing.T) {
		newUser, _ = userService.CreateUser(&model.NewUser{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@email.com",
		})

		assert.Equal(t, newUser.FirstName, "John")
		assert.Equal(t, newUser.LastName, "Doe")
	})

	t.Run("should_get_existing_user_by_id", func(t *testing.T) {
		id := newUser.ID.String()
		foundUser, err := userService.GetUser(id)

		assert.Equal(t, foundUser.FirstName, "John")
		assert.Equal(t, foundUser.LastName, "Doe")
		assert.Nil(t, err)
	})

	t.Run("should_get_service_name_called", func(t *testing.T) {
		assert.Equal(t, userService.ServiceName(), "UserServices")
	})
}
