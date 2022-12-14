package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
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
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get Abs path to avoid dir traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// no Abs path, respond with 400 bad request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend path with the path to the static dir
	path = filepath.Join(h.staticPath, path)

	// check whether file exist at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file doesn't exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// just another error we don't know about
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// server static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	// Initialize router
	r := mux.NewRouter()

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
	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logHandler := h.LoggingHandler{}
	spaHandler := spaHandler{staticPath: "client/build", indexPath: "index.html"}
	r.Handle("/graphql", playground.Handler("GraphQL", "/query"))
	r.Handle("/login", h.AddContext(ctx, logHandler.Logging(h.Login())))
	r.Handle("/query", h.AddContext(ctx, logHandler.Logging(h.Authenticate(srv))))

	r.PathPrefix("/").Handler(spaHandler)

	log.Infof("connect to http://localhost:%s/graphql for GraphQL playground", serverConfig.ServerPort)
	log.Fatal(server.ListenAndServe())
}
