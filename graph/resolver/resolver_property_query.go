package resolver

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// GetProperty is the resolver for the getProperty field.
func (r *queryResolver) GetProperty(ctx context.Context, id string) (*model.Property, error) {
	foundProperty, err := ctx.Value("propertyService").(*services.PropertyServices).GetProperty(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundProperty, nil
}
