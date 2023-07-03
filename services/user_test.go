package services

import (
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/util"
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

	twilioService := NewTwilioService(configuration.Twilio, queries)
	userService = NewUserService(queries, logger, &configuration.JwtConfig, twilioService)
	authService = NewAuthService(logger, &configuration.JwtConfig)
	amenityService = NewAmenityService(queries, logger, propertyService)
	propertyService = NewPropertyService(queries, logger, twilioService)
	unitService = NewUnitService(queries, logger)
	tenancyService = NewTenancyService(queries, logger)
	listingService = NewListingService(queries, logger)
	postaService = NewPostaService()

	// exit once done
	os.Exit(m.Run())
}

func Test_User_Services(t *testing.T) {
	var newUser *model.User
	var err error

	t.Run("should_create_new_user", func(t *testing.T) {
		newUser, err = userService.CreateUser(&model.NewUser{
			FirstName: "John",
			LastName:  "Doe",
			Email:     util.GenerateRandomEmail(),
			Avatar:    "https://avatar.jpg",
			Phone:     "+2541710073434",
		})

		assert.Nil(t, err)
		assert.Equal(t, newUser.FirstName, "John")
		assert.Equal(t, newUser.LastName, "Doe")
	})

	t.Run("should_sign_in_user", func(t *testing.T) {
		token, err := userService.SignIn(&model.NewUser{FirstName: "John", LastName: "Doe", Email: util.GenerateRandomEmail(), Phone: "+254190970274"})

		assert.Nil(t, err)
		assert.NotEqual(t, token, "")

		//	tokenValid, err := userService.ValidateToken(token)

		//assert.Nil(t, err)
		//assert.True(t, tokenValid.Valid)
	})

	t.Run("should_get_existing_user_by_id", func(t *testing.T) {
		foundUser, err := userService.FindById(newUser.ID)

		assert.Equal(t, foundUser.FirstName, "John")
		assert.Equal(t, foundUser.LastName, "Doe")
		assert.Nil(t, err)
	})

	t.Run("should_get_service_name_called", func(t *testing.T) {
		assert.Equal(t, userService.ServiceName(), "UserServices")
	})
}
