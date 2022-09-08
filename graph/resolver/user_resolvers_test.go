package resolver_test

import (
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/resolver"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func Test_Resolver_CreateUser(t *testing.T) {
	var _ = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))
}
