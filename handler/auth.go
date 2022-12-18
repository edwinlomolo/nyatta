package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

func Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		loginResponse := &model.LoginResponse{}
		var newUser *model.NewUser

		// Read incoming data from body request
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &newUser)

		// SignIn - authorize incoming user
		token, err := ctx.Value("userService").(*services.UserServices).SignIn(newUser)
		if err != nil {
			response := &model.Response{
				Code: http.StatusInternalServerError,
				Err:  err.Error(),
			}
			loginResponse.Response = response
			writeResponse(w, loginResponse, loginResponse.Code)
			return
		}

		response := &model.Response{
			Code: http.StatusCreated,
		}
		loginResponse.Response = response
		loginResponse.AccessToken = *token

		writeResponse(w, loginResponse, loginResponse.Code)
	})
}

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			userId string

			isAuthorized bool
		)

		// Authenticate - authenticate incoming request(s)
		ctx := r.Context()
		// Validate incoming Authorization header
		token, err := validateBearerAuthHeader(ctx, r)
		if err == nil {
			isAuthorized = true
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Grab user_id from jwt claims
				userIdBytes, _ := base64.StdEncoding.DecodeString(claims["id"].(string))
				userId = string(userIdBytes[:])
			}
		} else {
			// Failed authenticating Authorization header
			w.WriteHeader(http.StatusUnauthorized) // 401:Unauthorized
			w.Write([]byte("Unauthorized"))
			return
		}

		// Get remote user ip address from request
		userIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("Request Ip error: %v", err)
		}
		// Pass user info into current context
		ctx = context.WithValue(ctx, "ip", &userIp)
		ctx = context.WithValue(ctx, "userId", &userId)
		ctx = context.WithValue(ctx, "is_authorized", isAuthorized)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// validateBearerAuthHeader - validate incoming request Authorization header
func validateBearerAuthHeader(ctx context.Context, r *http.Request) (*jwt.Token, error) {
	var tokenString string
	// Grab header values from Authorization key and split into [`Bearer` and `{token}`] slice
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		return nil, errors.New(config.CredentialsError)
	}
	tokenString = auth[1]
	// Is token valid?
	token, err := ctx.Value("userService").(*services.UserServices).ValidateToken(&tokenString)
	return token, err
}

// writeResponse - feed response result into the final request response
func writeResponse(w http.ResponseWriter, response interface{}, code int) {
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
