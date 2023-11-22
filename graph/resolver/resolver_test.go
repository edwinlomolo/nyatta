package resolver

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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
	mailingService = services.NewMailingService(queries, configuration.Email, logger)
	twilioService = services.NewTwilioService(configuration.Twilio, queries, logger)
	userService = services.NewUserService(queries, logger, &configuration.JwtConfig, twilioService)
	propertyService = services.NewPropertyService(queries, logger, twilioService)
	amenityService = services.NewAmenityService(queries, logger, propertyService)
	unitService = services.NewUnitService(queries, logger)
	tenancyService = services.NewTenancyService(queries, logger)
	listingService = services.NewListingService(queries, logger)
	postaService = services.NewPostaService(logger)

	// Setup context
	ctx = context.Background()
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "amenityService", amenityService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "listingService", listingService)
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
