package main

import (
	"context"
	"net/http"

	"github.com/3dw1nM0535/nyatta/config"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Load env config(s)
	cfg, _ := config.LoadConfig(".")

	// Initialize service(s)
	ctx := context.Background()
	logger, _ := services.NewLogger(cfg)
	store, _ := config.OpenDB(cfg, logger)
	userService := services.NewUserService(store, logger, cfg)

	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "log", logger)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(resolver.New()))

	logHandler := h.LoggingHandler{}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/login", h.AddContext(ctx, logHandler.Logging(h.Login())))
	http.Handle("/query", h.AddContext(ctx, logHandler.Logging(h.Authenticate(srv))))

	logger.Debugf("connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	logger.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
