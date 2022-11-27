package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	SignJWT(user *model.User) (*string, error)
	ValidateJWT(token *string) (*jwt.Token, error)
}
