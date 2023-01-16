package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/3dw1nM0535/nyatta/config"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"

	"github.com/3dw1nM0535/nyatta/database"
	"github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	// Initialize router
	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)

	configuration := config.LoadConfig()

	serverConfig := configuration.Server

	// Initialize service(s)
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("%s: %v", database.DatabaseError, err)
	}
	queries := store.New(db)
	ctx := context.Background()
	logger := log.New()
	userService := services.NewUserService(queries, logger, &configuration.JwtConfig)
	propertyService := services.NewPropertyService(queries, logger)
	amenityService := services.NewAmenityService(queries, logger)
	unitService := services.NewUnitService(queries, logger)
	tenancyService := services.NewTenancyService(queries, logger)

	// Initialize context with values
	ctx = context.WithValue(ctx, "config", config.GetConfig())
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "amenityService", amenityService)
	ctx = context.WithValue(ctx, "unitService", unitService)
	ctx = context.WithValue(ctx, "tenancyService", tenancyService)
	ctx = context.WithValue(ctx, "log", logger)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.New()))

	logHandler := h.LoggingHandler{}
	r.Handle("/", playground.Handler("GraphQL", "/api"))
	r.Handle("/handshake", h.AddContext(ctx, logHandler.Logging(h.Handshake())))
	r.Handle("/api", h.AddContext(ctx, logHandler.Logging(h.Authenticate(srv))))

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverConfig.ServerPort),
		Handler: r,
	}

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", serverConfig.ServerPort)
	log.Fatal(s.ListenAndServe())
}
