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
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
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
	listingService  *services.ListingServices
	logger          *log.Logger
	configuration   *config.Configuration
	err             error
	db              *sql.DB
	postaService    *services.PostaServices
	twilioService   *services.TwilioServices
	mailingService  *services.MailingServices
)

// setup tests
func TestMain(m *testing.M) {
	// Load env variables
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Errorf("panic loading env: %v", err)
	}
	configuration = config.LoadConfig()

	// Initialize db
	db, err = database.InitDB("../../database/migration")
	if err != nil {
		log.Fatalf("%s: %v", database.DatabaseError, err)
	}

	// Logger
	logger = log.New()

	// SQL queries
	queries := sqlStore.New(db)

	// Setup services
	twilioService = services.NewTwilioService(configuration.Twilio, queries)
	userService = services.NewUserService(queries, logger, &configuration.JwtConfig, twilioService)
	propertyService = services.NewPropertyService(queries, logger, twilioService)
	amenityService = services.NewAmenityService(queries, logger, propertyService)
	unitService = services.NewUnitService(queries, logger)
	tenancyService = services.NewTenancyService(queries, logger)
	listingService = services.NewListingService(queries, logger)
	postaService = services.NewPostaService()
	mailingService = services.NewMailingService(queries, configuration.Email)

	// Setup context
	ctx = context.Background()
	ctx = context.WithValue(ctx, "configuration", configuration)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "amenityService", amenityService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "listingService", listingService)
	ctx = context.WithValue(ctx, "log", logger)
	ctx = context.WithValue(ctx, "store", db)
	ctx = context.WithValue(ctx, "postaService", postaService)
	ctx = context.WithValue(ctx, "twilioService", twilioService)
	ctx = context.WithValue(ctx, "mailingService", mailingService)

	// Run test
	exitCode := m.Run()

	// Teardown test
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Errorf("panic loading postgres driver instance: %v", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance("file://../../database/migration", "postgres", driver)
	if err != nil {
		log.Errorf("panic tearing down test: %v", err)
	}
	if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
		log.Errorf("panic migrator err: %v", err)
	}

	os.Exit(exitCode)
}

// TODO redudant test
func Test_unauthed_graphql_request(t *testing.T) {
	var signIn struct {
		SignIn struct {
			Token string
		}
	}
	var srv = client.New(h.AddContext(context.Background(), h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))))

	t.Run("should_not_next_unauthed_graphql_request", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { signIn (input: { first_name: %q, last_name: %q, email: %q, avatar: %q }) { token } }`, "Jane", "Doe", util.GenerateRandomEmail(), "https://avatar.jpg")

		err := srv.Post(query, &signIn)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "http 401: {\"Unauthorized\":true}")
	})
}
