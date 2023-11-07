package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/3dw1nM0535/nyatta/config"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"

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

	// TODO provide mpesa callback for pay services
	mpesaService := services.NewMpesaService(configuration.Mpesa, logger)
	t := time.Now().Format("20060102150405")
	dataToEncode := fmt.Sprintf("%d%s%s", 174379, configuration.Mpesa.PassKey, t)
	res, err := mpesaService.StkPush(services.LipaNaMpesaPayload{
		BusinessShortCode: 174379,
		Password:          base64.StdEncoding.EncodeToString([]byte(dataToEncode)),
		Timestamp:         t,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            1,
		PartyA:            254792921440,
		PartyB:            174379,
		PhoneNumber:       254792921440,
		CallBackURL:       "https://d86d-102-217-127-1.ngrok.io/mpesa/charge",
		AccountReference:  "CompanyXLTD",
		TransactionDesc:   "Landlord subscription",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	mailingService := services.NewMailingService(queries, configuration.Email, logger)
	twilioService := services.NewTwilioService(configuration.Twilio, queries, logger)
	userService := services.NewUserService(queries, logger, configuration.Server.ServerEnv, &configuration.JwtConfig, twilioService, mailingService.SendEmail)
	propertyService := services.NewPropertyService(queries, configuration.Server.ServerEnv, logger, twilioService, mailingService.SendEmail)
	amenityService := services.NewAmenityService(queries, logger, propertyService)
	unitService := services.NewUnitService(queries, logger)
	tenancyService := services.NewTenancyService(queries, logger)
	listingService := services.NewListingService(queries, logger)
	postaService := services.NewPostaService(logger)
	awsService := services.NewAwsService(configuration.Aws, logger)

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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.New()))

	logHandler := h.LoggingHandler{}
	r.Handle("/", playground.Handler("GraphQL", "/api"))
	r.Handle("/mpesa/charge", h.AddContext(ctx, logHandler.Logging(h.MpesaChargeCallback())))
	r.Handle("/api", h.AddContext(ctx, logHandler.Logging(h.Authenticate(srv))))

	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", serverConfig.ServerPort),
		Handler: r,
	}

	logrus.Infof("Server Info: OK")
	logrus.Fatal(s.ListenAndServe())
}
