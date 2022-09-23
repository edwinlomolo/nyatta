package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
)

func Test_Auth_Handler(t *testing.T) {
	// Load env config(s)
	cfg, _ := nyatta_context.LoadConfig("..")

	// Initialize service(s)
	ctx := context.Background()
	logger, _ := services.NewLogger(cfg)
	store, _ := nyatta_context.OpenDB(cfg, logger)
	userService := services.NewUserService(store, logger, cfg)

	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "log", logger)

	var creds struct {
		AccessToken string `json:"access_token"`
		Code        int    `json:"code"`
	}

	var jsonStr = []byte(fmt.Sprintf(`{"first_name": "%s", "last_name": "%s", "email": "%s"}`, "john", "doe", util.GenerateRandomEmail()))

	t.Run("should_login_ok_by_login_handler", func(t *testing.T) {
		srv := httptest.NewServer(AddContext(ctx, Login()))

		defer srv.Close()

		url := fmt.Sprintf("%s/login", srv.URL)
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))

		client := srv.Client()
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}

		json.Unmarshal(data, &creds)
		if creds.AccessToken == "" {
			t.Errorf("expected token to be provided, got %s", creds.AccessToken)
		}
		if creds.Code != 201 {
			t.Errorf("expected 201 created code, got %d", creds.Code)
		}
	})

	t.Run("should_drop_any_unauthed_request_successfully", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w)
		})

		srv := httptest.NewServer(AddContext(ctx, Authenticate(handler)))
		defer srv.Close()

		res, err := http.Get(srv.URL)
		if err != nil {
			t.Errorf("expected err to be nil, got: %v", err)
		}

		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Errorf("expected err to be nil, got %v", err)
		}

		if string(data) != "Unauthorized" {
			t.Errorf("got %s expected Unauthorized", string(data))
		}

		if res.Status != "401 Unauthorized" {
			t.Errorf("expected 401 Unauthorized got %v", res.Status)
		}
	})
	t.Run("should_authenticate_any_authed_request_successfully", func(t *testing.T) {
		tokenString := fmt.Sprintf("Bearer %s", creds.AccessToken)
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world")
		})

		srv := httptest.NewServer(AddContext(ctx, Authenticate(handler)))
		defer srv.Close()

		req, _ := http.NewRequest(http.MethodPost, srv.URL, bytes.NewBuffer(jsonStr))
		req.Header.Add("Authorization", tokenString)

		client := srv.Client()
		res, err := client.Do(req)
		if err != nil {
			t.Errorf("expected err to be nil got %v", err)
		}

		if res.Status != "200 OK" {
			t.Errorf("expected 200 OK status got %v", res.Status)
		}

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected err to be nil got %v", err)
		}

		if string(data) != "Hello, world" {
			t.Errorf("expected Hello, world got %s", string(data))
		}

	})
}
