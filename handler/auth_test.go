package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var ctx context.Context

func TestMain(m *testing.M) {
	// Load env config(s)
	logger := log.New()
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		logger.Errorf("panic loading env: %v", err)
	}
	cfg := config.LoadConfig()

	// Initialize service(s)
	ctx = context.Background()
	db, err := database.InitDB("../database/migration")
	if err != nil {
		log.Fatalf("%s: %v", database.DatabaseError, err)
	}
	queries := sqlStore.New(db)
	userService := services.NewUserService(queries, logger, &cfg.JwtConfig)
	propertyService := services.NewPropertyService(queries, logger)
	unitService := services.NewUnitService(queries, logger)
	tenancyService := services.NewTenancyService(queries, logger)

	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "log", logger)

	// exit once done
	os.Exit(m.Run())
}

func Test_Auth_Handler(t *testing.T) {
	var creds struct {
		AccessToken string `json:"access_token"`
		Code        int    `json:"code"`
	}

	var jsonStr = []byte(fmt.Sprintf(`{"first_name": %q, "last_name": %q, "email": %q, "avatar": %q}`, "john", "doe", util.GenerateRandomEmail(), "https://avatar.jpg"))

	t.Run("should_login_ok_by_login_handler", func(t *testing.T) {
		srv := httptest.NewServer(AddContext(ctx, Handshake()))

		defer srv.Close()

		url := fmt.Sprintf("%s/handshake", srv.URL)
		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))

		client := srv.Client()
		res, err := client.Do(req)
		assert.Nil(t, err)

		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)

		json.Unmarshal(data, &creds)
		assert.NotEmpty(t, creds.AccessToken)
		assert.Equal(t, creds.Code, 200)
	})

	t.Run("handshake_endpoint_should_only_support_POST_method", func(t *testing.T) {
		srv := httptest.NewServer(AddContext(ctx, Handshake()))
		defer srv.Close()

		url := fmt.Sprintf("%s/handshake", srv.URL)
		req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonStr))

		client := srv.Client()
		res, err := client.Do(req)
		assert.Nil(t, err)

		defer res.Body.Close()

		data, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, string(data), "Only POST method supported\n")
		assert.Equal(t, res.StatusCode, http.StatusMethodNotAllowed)
	})
}
