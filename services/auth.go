package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/interfaces"
	jwt "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

// AuthServices - represent authentication service
type AuthServices struct {
	log       *log.Logger
	secret    *string
	expiresIn *time.Duration
}

// _ - AuthServices{} implements AuthService
var _ interfaces.AuthService = &AuthServices{}

// NewAuthService - factory for auth services
func NewAuthService(logger *log.Logger, config *config.Jwt) *AuthServices {
	return &AuthServices{logger, &config.JWT.Secret, &config.JWT.Expires}
}

// SignJWT - signin user and return jwt token
func (a *AuthServices) SignJWT(ctx context.Context, user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":         "Nyatta",
		"created_at":  user.CreatedAt,
		"is_landlord": user.IsLandlord,
		"user_phone":  user.Phone,
		"id":          base64.StdEncoding.EncodeToString([]byte(user.ID.String())),
		"exp":         time.Now().Add(time.Second * *a.expiresIn).Unix(),
	})

	tokenString, err := token.SignedString([]byte(*a.secret))
	if err != nil {
		a.log.Errorf("%s: %v", a.ServiceName(), err)
		return nil, err
	}

	return &tokenString, nil
}

// ValidateJWT - validate jwt token
func (a *AuthServices) ValidateJWT(ctx context.Context, tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			a.log.Errorf("%s: %v", a.ServiceName(), fmt.Errorf("%s: %v", config.InvalidTokenSigningAlgorithm.Error(), token.Header["alg"]))
			return nil, fmt.Errorf("%s: %v", config.InvalidTokenSigningAlgorithm.Error(), token.Header["alg"])
		}

		return []byte(*a.secret), nil
	})

	if err != nil {
		a.log.Errorf("%s: %v", a.ServiceName(), fmt.Errorf("%s: %v", config.TokenParsing.Error(), err))
		return nil, fmt.Errorf("%s: %v", config.TokenParsing.Error(), err)
	}

	return token, nil
}

func (a AuthServices) ServiceName() string {
	return "AuthServices"
}
