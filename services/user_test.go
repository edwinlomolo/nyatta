package services

import (
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var (
	userService     *UserServices
	authService     *AuthServices
	propertyService *PropertyServices
	listingService  *ListingServices
	amenityService  *AmenityServices
	unitService     *UnitServices
	tenancyService  *TenancyServices
	queries         *sqlStore.Queries
	configuration   *config.Configuration
	postaService    *PostaServices
	twilioService   *TwilioServices
	mailingService  *MailingServices
)

func TestMain(m *testing.M) {
	// setup test
	logger := log.New()
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		logger.Errorf("panic loading env: %v", err)
	}
	configuration := config.LoadConfig()
	db, err := database.InitDB("../database/migration")
	if err != nil {
		log.Fatalf("%s: %v", database.DatabaseError, err)
	}
	queries = sqlStore.New(db)

	twilioService = NewTwilioService(configuration.Twilio, queries, logger)
	mailingService = NewMailingService(queries, configuration.Email, logger)
	userService = NewUserService(queries, logger, &configuration.JwtConfig, twilioService)
	authService = NewAuthService(logger, &configuration.JwtConfig)
	amenityService = NewAmenityService(queries, logger)
	propertyService = NewPropertyService(queries, logger, twilioService)
	unitService = NewUnitService(queries, logger)
	tenancyService = NewTenancyService(queries, logger)
	listingService = NewListingService(queries, logger)
	postaService = NewPostaService(logger)

	// exit once done
	os.Exit(m.Run())
}

func Test_User_Services(t *testing.T) {
	t.Run("should_create_new_user", func(t *testing.T) {
	})

	t.Run("should_sign_in_user", func(t *testing.T) {
	})

	t.Run("should_get_existing_user_by_id", func(t *testing.T) {
	})

	t.Run("should_get_service_name_called", func(t *testing.T) {
		assert.Equal(t, userService.ServiceName(), "UserServices")
	})
}
