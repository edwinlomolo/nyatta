package interfaces

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserService interface {
	SignIn(ctx context.Context, user *model.NewUser) (*model.SignIn, error)
	UpdateUserInfo(ctx context.Context, id uuid.UUID, firstName, lastName, avatar string) (*model.User, error)
	GetUserAvatar(ctx context.Context, id uuid.UUID) (*model.AnyUpload, error)
	ValidateToken(ctx context.Context, token *string) (*jwt.Token, error)
	ServiceName() string
	FindUserByPhone(ctx context.Context, phone string) (*model.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*model.User, error)
}
