package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/3dw1nM0535/nyatta/graph/model"
)

func Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loginResponse := &model.LoginResponse{}
		_, err := validateBasicAuthHeader(r)
		if err != nil {
			response := &model.Response{
				Code: http.StatusInternalServerError,
				Err:  err.Error(),
			}
			loginResponse.Response = response
			writeResponse(w, loginResponse, loginResponse.Code)
		}
		response := &model.Response{
			Code: http.StatusOK,
		}
		loginResponse.Response = response
		loginResponse.AccessToken = "token"
		writeResponse(w, loginResponse, loginResponse.Code)
	})
}

func validateBasicAuthHeader(r *http.Request) (*model.UserCredentials, error) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return nil, fmt.Errorf("%d: credentials error", http.StatusBadRequest)
	}
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return nil, fmt.Errorf("%d: credentials error", http.StatusBadRequest)
	}
	return &model.UserCredentials{
		Email:    pair[0],
		Password: pair[1],
	}, nil
}

func writeResponse(w http.ResponseWriter, response interface{}, code int) {
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
