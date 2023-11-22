package handler

import (
	"context"
	"encoding/base64"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/services"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			userId       string
			isAuthorized bool
			isLandlord   bool
			phone        string
		)

		// Authenticate - authenticate incoming request(s)
		ctx := r.Context()
		logger := ctx.Value("log").(*logrus.Logger)

		// Validate incoming Authorization header
		token, err := validateBearerAuthHeader(ctx, r)
		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Grab user_id from jwt claims
				userIdBytes, _ := base64.StdEncoding.DecodeString(claims["id"].(string))
				userId = string(userIdBytes[:])
				phone = claims["user_phone"].(string)
				isLandlord = claims["is_landlord"].(bool)
				isAuthorized = true
			}
		} else {
			if err != nil {
				logger.Errorf("%s:%v", "AuthMiddlewareTokenValidationError", err)
				//http.Error(w, err.Error(), http.StatusNotAuthorized)
			}
		}

		// Get remote user ip address from request
		userIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("Request Ip error: %v", err)
		}
		// Pass user info into current context
		ctx = context.WithValue(ctx, "ip", userIp)
		ctx = context.WithValue(ctx, "userId", userId)
		ctx = context.WithValue(ctx, "is_landlord", isLandlord)
		ctx = context.WithValue(ctx, "is_authorized", isAuthorized)
		ctx = context.WithValue(ctx, "phone", phone)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// validateBearerAuthHeader - validate incoming request Authorization header
func validateBearerAuthHeader(ctx context.Context, r *http.Request) (*jwt.Token, error) {
	var tokenString string
	// Grab header values from Authorization key and split into [`Bearer` and `{token}`] slice
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		return nil, config.CredentialsError
	}
	tokenString = auth[1]
	// Is token valid?
	token, err := ctx.Value("userService").(*services.UserServices).ValidateToken(ctx, &tokenString)
	return token, err
}
