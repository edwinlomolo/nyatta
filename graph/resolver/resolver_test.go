package resolver

import (
	"context"
	"database/sql"
	"fmt"
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
	"github.com/stretchr/testify/assert"
)

var (
	ctx             context.Context
	userService     *services.UserServices
	propertyService *services.PropertyServices
	amenityService  *services.AmenityServices
	unitService     *services.UnitServices
	tenancyService  *services.TenancyServices
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
	amenityService = services.NewAmenityService(queries, logger)
	unitService = services.NewUnitService(queries, logger)
	tenancyService = services.NewTenancyService(queries, logger)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "config", configuration)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "amenityService", amenityService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "log", logger)

	os.Exit(m.Run())
}
func Test_unauthed_graphql_request(t *testing.T) {
	var signIn struct {
		SignIn struct {
			Token string
		}
	}
	var srv = client.New(h.AddContext(context.Background(), h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))))

	t.Run("should_not_next_unauthed_graphql_request", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { signIn (input: { first_name: "%s", last_name: "%s", email: "%s" }) { token } }`, "Jane", "Doe", util.GenerateRandomEmail())

		err := srv.Post(query, &signIn)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "http 401: Unauthorized")
	})
}
