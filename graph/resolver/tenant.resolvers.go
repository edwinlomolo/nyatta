package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// User is the resolver for the user field.
func (r *tenantResolver) User(ctx context.Context, obj *model.Tenant) (*model.User, error) {
	user, err := ctx.Value("userService").(services.UserService).GetUser(ctx, obj.UserID)
	if err != nil {
		return nil, err
	}

	return user, err
}

// Tenant returns generated.TenantResolver implementation.
func (r *Resolver) Tenant() generated.TenantResolver { return &tenantResolver{r} }

type tenantResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *tenantResolver) Unit(ctx context.Context, obj *model.Tenant) (*model.Unit, error) {
	unit, err := ctx.Value("unitService").(services.UnitService).GetUnit(ctx, obj.UnitID)
	if err != nil {
		return nil, err
	}

	return unit, err
}
