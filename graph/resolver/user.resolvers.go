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

// Avatar is the resolver for the avatar field.
func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented: Avatar - avatar"))
}

// Properties is the resolver for the properties field.
func (r *userResolver) Properties(ctx context.Context, obj *model.User) ([]*model.Property, error) {
	userProperties, err := ctx.Value("propertyService").(*services.PropertyServices).PropertiesCreatedBy(obj.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return userProperties, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
