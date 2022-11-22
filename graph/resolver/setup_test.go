package resolver

import (
	"bytes"
	"context"
	"database/sql"
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
	"github.com/3dw1nM0535/nyatta/graph/generated"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	ctx             context.Context
	userService     *services.UserServices
	propertyService *services.PropertyServices
	logger          *log.Logger
	configuration   *config.Configuration
	err             error
	db              *sql.DB
)

// setup tests
func TestMain(m *testing.M) {
	logger := log.New()
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Errorf("panic loading env: %v", err)
	}
	configuration = config.LoadConfig()
	db, err = database.InitDB()
	if err != nil {
		log.Fatalf("%s: %v", database.DatabaseError, err)
	}
	queries := sqlStore.New(db)

	userService = services.NewUserService(queries, logger, &configuration.JwtConfig)
	propertyService = services.NewPropertyService(queries, logger)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "config", configuration)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "log", logger)

	os.Exit(m.Run())
}

// makeLoginUser - return authed user
func makeLoginUser() string {
	var creds struct {
		AccessToken string `json:"access_token"`
		Code        int    `json:"code"`
	}
	var jsonStr = []byte(fmt.Sprintf(`{"first_name": "%s", "last_name": "%s", "email": "%s"}`, "john", "doe", util.GenerateRandomEmail()))

	httpServer := httptest.NewServer(h.AddContext(ctx, h.Login()))
	defer httpServer.Close()

	url := fmt.Sprintf("%s/login", httpServer.URL)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))

	c := httpServer.Client()
	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(data, &creds)
	return creds.AccessToken
}

// makeAuthedGqlServer - return authenticated graphql client
func makeAuthedGqlServer(authenticate bool, ctx context.Context) *client.Client {
	var srv *client.Client
	if !authenticate {
		// unauthed client
		srv = client.New(h.AddContext(ctx, h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))))
		return srv
	}
	// authed user
	tokenString := makeLoginUser()
	// authed client
	srv = client.New(h.AddContext(ctx, h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))), client.AddHeader("Authorization", fmt.Sprintf("Bearer %s", tokenString)))

	return srv
}
