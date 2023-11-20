package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserService interface {
	SignIn(user *model.NewUser) (*model.SignIn, error)
	UpdateUserInfo(id uuid.UUID, phone, firstName, lastName, avatar string) (*model.User, error)
	GetUserAvatar(id uuid.UUID) (*model.AnyUpload, error)
	ValidateToken(token *string) (*jwt.Token, error)
	ServiceName() string
	FindUserByPhone(phone string) (*model.User, error)
	GetUser(id uuid.UUID) (*model.User, error)
}
