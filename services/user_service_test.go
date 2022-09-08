package services

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/stretchr/testify/assert"
)

var (
	userService *UserServices
)

func Test_CreateUser(t *testing.T) {
	newUser := &model.NewUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
	}
	assert.EqualValues(t, newUser.FirstName, "John")
}
