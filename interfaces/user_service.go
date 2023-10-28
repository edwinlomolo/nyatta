package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	UpdateUser(input *model.UpdateUserInput) (*model.User, error)
	SignIn(user *model.NewUser) (*model.SignIn, error)
	ValidateToken(token *string) (*jwt.Token, error)
	ServiceName() string
	FindUserByPhone(phone string) (*model.User, error)
	OnboardUser(email string, onboarding bool) (*model.User, error)
}
