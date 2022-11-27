package services

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// AuthServices - represent authentication service
type AuthServices struct {
	log       *log.Logger
	secret    *string
	expiresIn *time.Duration
}

//_ - AuthServices{} implements AuthService
var _ interfaces.AuthService = &AuthServices{}

// NewAuthService - factory for auth services
func NewAuthService(logger *log.Logger, config *config.Jwt) *AuthServices {
	return &AuthServices{logger, &config.JWT.Secret, &config.JWT.Expires}
}

// SignJWT - signin user and return jwt token
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

// ValidateJWT - validate jwt token
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
