package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	CreateUser(user *model.NewUser) (*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	UpdateUser(input *model.UpdateUserInput) (*model.User, error)
	SignIn(user *model.NewUser) (*string, error)
	ValidateToken(token *string) (*jwt.Token, error)
	ServiceName() string
	FindUserByPhone(phone string) (*model.User, error)
}
