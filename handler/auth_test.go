package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/model"
)

func Test_Auth_Handler(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.SetBasicAuth("admin", "1234")
		loginResponse := &model.LoginResponse{}
		_, err := validateBasicAuthHeader(r)
		if err != nil {
			response := &model.Response{
				Code: http.StatusBadRequest,
				Err:  err.Error(),
			}
			loginResponse.Response = response
			writeResponse(w, loginResponse, loginResponse.Code)
			return
		}
		response := &model.Response{
			Code: http.StatusOK,
		}
		loginResponse.Response = response
		loginResponse.AccessToken = "token"
		writeResponse(w, loginResponse, loginResponse.Code)
	}))
	defer srv.Close()

	c := srv.Client()

	res, err := c.Post(fmt.Sprintf("%s/login", srv.URL), "application/json", nil)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	var creds struct {
		AccessToken string `json:"access_token"`
	}
	json.Unmarshal(data, &creds)
	if creds.AccessToken == "" {
		t.Errorf("expected token to be provided, got %s", creds.AccessToken)
	}
}
