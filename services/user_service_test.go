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

func Test_WhichService(t *testing.T) {
	serviceName := userService.ServiceName()

	assert.Equal(t, serviceName, "UserServices")
}

func Test_CreateUser(t *testing.T) {
	newUser, _ := userService.CreateUser(&model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
	})

	assert.EqualValues(t, newUser.FirstName, "")
}

func Test_GetUser(t *testing.T) {
	foundUser, _ := userService.GetUser("1")

	assert.EqualValues(t, foundUser.ID, "")
}
