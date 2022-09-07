package main

import (
	"log"
	"net/http"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	service "github.com/3dw1nM0535/nyatta/services"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	defaultPort = "8080"
)

func main() {
	// Load env config(s)
	cfg, _ := nyatta_context.LoadConfig(".")

	// Initialize service(s)
	logger, _ := service.NewLogger(cfg)
	store, _ := nyatta_context.OpenDB(cfg, logger)
	userService := service.NewUserService(store, logger)

	port := cfg.Port
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UserService: userService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
