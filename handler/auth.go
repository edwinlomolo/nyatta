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

func Handshake() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		loginResponse := &model.LoginResponse{}
		var newUser *model.NewUser

		// This handler should only support POST
		if r.Method != http.MethodPost {
			http.Error(w, errors.New("Only POST method supported").Error(), http.StatusMethodNotAllowed)
			return
		}
		// Read incoming data from body request
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.Unmarshal(reqBody, &newUser)

		// SignIn - authorize incoming user
		token, err := ctx.Value("userService").(*services.UserServices).SignIn(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {

			response := &model.Response{
				Code: http.StatusOK,
			}
			loginResponse.Response = response
			loginResponse.AccessToken = *token

			writeResponse(w, loginResponse, loginResponse.Code)
			return
		}
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
			// Failed
			jsonResponse, err := json.Marshal(struct{ Unauthorized bool }{Unauthorized: true})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
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
		return nil, config.CredentialsError
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
