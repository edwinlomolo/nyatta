package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate --verbose
import (
	"github.com/3dw1nM0535/nyatta/graph/generated"
)

type Resolver struct{}

func New() generated.Config {
	c := generated.Config{Resolvers: &Resolver{}}
	return c
}
