package main

import (
	"context"
	"net/http"

	"github.com/3dw1nM0535/nyatta/config"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"

	"github.com/3dw1nM0535/nyatta/database"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	log "github.com/sirupsen/logrus"
)

func main() {
	configuration := config.LoadConfig()

	serverConfig := configuration.Server

	// Initialize service(s)
	ctx := context.Background()
	logger := log.New()
	store, _ := database.InitDB()
	userService := services.NewUserService(store, logger, &configuration.JwtConfig)

	ctx = context.WithValue(ctx, "config", config.GetConfig())
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "log", logger)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.New()))

	logHandler := h.LoggingHandler{}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/login", h.AddContext(ctx, logHandler.Logging(h.Login())))
	http.Handle("/query", h.AddContext(ctx, logHandler.Logging(h.Authenticate(srv))))

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", serverConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, nil))
}
