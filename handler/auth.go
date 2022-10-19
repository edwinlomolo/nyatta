package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	jwt "github.com/dgrijalva/jwt-go"
)

func Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		loginResponse := &model.LoginResponse{}
		var newUser *model.NewUser

		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &newUser)
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
			userId       string
			isAuthorized bool
		)

		ctx := r.Context()
		token, err := validateBearerAuthHeader(ctx, r)
		if err == nil {
			isAuthorized = true
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userIdBytes, _ := base64.StdEncoding.DecodeString(claims["id"].(string))
				userId = string(userIdBytes[:])
			}
		} else {
			// TODO use logger
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		userIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("Request Ip error: %v", err)
		}
		ctx = context.WithValue(ctx, "ip", &userIp)
		ctx = context.WithValue(ctx, "userId", &userId)
		ctx = context.WithValue(ctx, "is_authorized", isAuthorized)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateBearerAuthHeader(ctx context.Context, r *http.Request) (*jwt.Token, error) {
	var tokenString string
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		return nil, errors.New(config.CredentialsError)
	}
	tokenString = auth[1]
	token, err := ctx.Value("userService").(*services.UserServices).ValidateToken(&tokenString)
	return token, err
}

func writeResponse(w http.ResponseWriter, response interface{}, code int) {
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
