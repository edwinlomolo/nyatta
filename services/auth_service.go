package services

import (
	"encoding/base64"
	"fmt"
	"time"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type AuthService interface {
	SignJwt(user *model.User) (string, error)
}

type AuthServices struct {
	log       *zap.SugaredLogger
	secret    *string
	expiresIn *time.Duration
}

func NewAuthService(logger *zap.SugaredLogger, config *nyatta_context.Config) *AuthServices {
	return &AuthServices{logger, &config.JWTSecret, &config.JWTExpiration}
}

func (a *AuthServices) SignJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":        "Nyatta",
		"created_at": user.CreatedAt,
		"id":         base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"exp":        time.Now().Add(time.Second * *a.expiresIn).Unix(),
	})

	tokenString, err := token.SignedString([]byte(*a.secret))
	return &tokenString, err
}

func (a *AuthServices) ValidateJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(*a.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unexpected error while parsing token: %v", err)
	}
	return token, nil
}
