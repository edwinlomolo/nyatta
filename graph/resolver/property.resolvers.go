package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// Owner is the resolver for the owner field.
func (r *propertyResolver) Owner(ctx context.Context, obj *model.Property) (*model.User, error) {
	foundOwner, err := ctx.Value("userService").(*services.UserServices).FindById(obj.CreatedBy)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundOwner, nil
}

// Property returns generated.PropertyResolver implementation.
func (r *Resolver) Property() generated.PropertyResolver { return &propertyResolver{r} }

type propertyResolver struct{ *Resolver }
