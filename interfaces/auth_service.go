package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	SignJWT(ctx context.Context, user *model.User) (*string, error)
	ValidateJWT(ctx context.Context, token *string) (*jwt.Token, error)
	ServiceName() string
}
