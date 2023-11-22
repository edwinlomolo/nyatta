package handler

import (
	"context"
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

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

	mailingService := services.NewMailingService(queries, cfg.Email, logger)
	twilioService := services.NewTwilioService(cfg.Twilio, queries, logger)
	userService := services.NewUserService(queries, logger, &cfg.JwtConfig, twilioService)
	propertyService := services.NewPropertyService(queries, logger, twilioService)
	unitService := services.NewUnitService(queries, logger)
	tenancyService := services.NewTenancyService(queries, logger)

	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "mailingService", mailingService)
	ctx = context.WithValue(ctx, "log", logger)

	// exit once done
	os.Exit(m.Run())
}

func Test_Auth_Handler(t *testing.T) {}
