package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/3dw1nM0535/nyatta/database"
	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	sentryLogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	// Initialize router
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)

	configuration := config.LoadConfig()

	serverConfig := configuration.Server

	// Start service(s) initialization
	db, err := database.InitDB("./database/migration")
	if err != nil {
		logrus.Fatalf("%s: %v", database.DatabaseError, err)
	}
	queries := store.New(db)
	ctx := context.Background()
	// Logging
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout

	if serverConfig.ServerEnv == "production" || serverConfig.ServerEnv == "staging" {
		// Error level to extract from logging
		sentryLevels := []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}
		sentryHook, err := sentryLogrus.New(sentryLevels, sentry.ClientOptions{
			Dsn:              configuration.SentryConfig.Dsn,
			AttachStacktrace: true,
		})
		if err != nil {
			logrus.Fatalf("Failed to initialize sentry logrus hook: %v", err)
		} else if err == nil {
			logger.AddHook(sentryHook)
			logrus.Infoln("Sentry logging enabled")
		}
		defer sentryHook.Flush(5 * time.Second)
		// Flush before calling os.Exit(1) on logger
		logrus.RegisterExitHandler(func() { sentryHook.Flush(5 * time.Second) })
	}

	mailingService := services.NewMailingService(queries, configuration.Email, logger)
	twilioService := services.NewTwilioService(configuration.Twilio, queries, logger)
	userService := services.NewUserService(queries, logger, &configuration.JwtConfig, twilioService)
	propertyService := services.NewPropertyService(queries, logger, twilioService)
	amenityService := services.NewAmenityService(queries, logger)
	unitService := services.NewUnitService(queries, logger)
	tenancyService := services.NewTenancyService(queries, logger)
	listingService := services.NewListingService(queries, logger)
	postaService := services.NewPostaService(logger)
	awsService := services.NewAwsService(configuration.Aws, logger)
	paystackService := services.NewPaystackService(configuration.Paystack, serverConfig.ServerEnv, logger, queries)
	equityBankService := services.NewEquityBankService(logger, serverConfig.ServerEnv, configuration.EquityBank)

	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "amenityService", amenityService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "listingService", listingService)
	ctx = context.WithValue(ctx, "postaService", postaService)
	ctx = context.WithValue(ctx, "awsService", awsService)
	ctx = context.WithValue(ctx, "log", logger)
	ctx = context.WithValue(ctx, "twilioService", twilioService)
	ctx = context.WithValue(ctx, "mailingService", mailingService)
	ctx = context.WithValue(ctx, "sqlStore", queries)
	ctx = context.WithValue(ctx, "paystackService", paystackService)
	ctx = context.WithValue(ctx, "equityBankService", equityBankService)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.New()))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logger.Errorf("%s:%v", "SetRecoverFuncGqlError", err)
		return gqlerror.Errorf("Internal server error")
	})

	logHandler := h.LoggingHandler{}
	r.Handle("/", playground.Handler("GraphQL", "/api"))
	r.Handle("/api", h.AddContext(ctx, logHandler.Logging(h.Authenticate(srv))))
	r.Method("POST", "/paystack/webhook", h.AddContext(ctx, logHandler.Logging(h.MpesaChargeCallback())))
	r.Method("POST", "/upload", h.AddContext(ctx, logHandler.Logging(h.Authenticate(h.UploadHandler()))))

	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", serverConfig.ServerPort),
		Handler: r,
	}

	logrus.Infof("Server Info: OK")
	logrus.Fatal(s.ListenAndServe())
}
