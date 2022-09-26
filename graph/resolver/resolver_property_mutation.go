package resolver

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/model"
)

// CreateProperty is the resolver for the createProperty field.
func (r *mutationResolver) CreateProperty(ctx context.Context, input model.NewProperty) (*model.Property, error) {
	panic(fmt.Errorf("not implemented: CreateProperty - createProperty"))
}
