package services

import (
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
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
	userService = NewUserService(queries, logger, configuration.Server.ServerEnv, &configuration.JwtConfig, twilioService, mailingService.SendEmail)
	authService = NewAuthService(logger, &configuration.JwtConfig)
	amenityService = NewAmenityService(queries, logger, propertyService)
	propertyService = NewPropertyService(queries, configuration.Server.ServerEnv, logger, twilioService, mailingService.SendEmail)
	unitService = NewUnitService(queries, logger)
	tenancyService = NewTenancyService(queries, logger)
	listingService = NewListingService(queries, logger)
	postaService = NewPostaService(logger)

	// exit once done
	os.Exit(m.Run())
}

func Test_User_Services(t *testing.T) {
	var newUser *model.User
	var err error

	t.Run("should_create_new_user", func(t *testing.T) {
		newUser, err = userService.FindUserByPhone("+2541710073434")

		assert.Nil(t, err)
	})

	t.Run("should_sign_in_user", func(t *testing.T) {
		res, err := userService.SignIn(&model.NewUser{Phone: "+254190970274"})

		assert.Nil(t, err)
		assert.NotEmpty(t, res.Token, "")
	})

	t.Run("should_get_existing_user_by_id", func(t *testing.T) {
		_, err := userService.FindById(newUser.ID)

		assert.Nil(t, err)
	})

	t.Run("should_get_service_name_called", func(t *testing.T) {
		assert.Equal(t, userService.ServiceName(), "UserServices")
	})
}
