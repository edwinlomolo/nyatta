package resolver_test

import (
	"context"
	"testing"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ctx           context.Context
	userService   *services.UserServices
	logger        *zap.SugaredLogger
	store         *gorm.DB
	cfg           *nyatta_context.Config
	configuration *nyatta_context.Config
)

func init() {
	configuration, _ = nyatta_context.LoadConfig("../../")
	logger, _ = services.NewLogger(configuration)
	store, _ = nyatta_context.OpenDB(configuration, logger)
	userService = services.NewUserService(store, logger)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "log", logger)
}

func Test_Resolver_CreateUser(t *testing.T) {
	var _ = client.New(h.AddContext(ctx, handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))))
}
