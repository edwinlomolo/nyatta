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

// Amenities is the resolver for the amenities field.
func (r *propertyResolver) Amenities(ctx context.Context, obj *model.Property) ([]*model.Amenity, error) {
	foundAmenities, err := ctx.Value("amenityService").(*services.AmenityServices).PropertyAmenities(obj.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundAmenities, nil
}

// Units is the resolver for the units field.
func (r *propertyResolver) Units(ctx context.Context, obj *model.Property) ([]*model.PropertyUnit, error) {
	foundUnits, err := ctx.Value("propertyService").(*services.PropertyServices).GetPropertyUnits(obj.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundUnits, nil
}

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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *propertyResolver) CreatedBy(ctx context.Context, obj *model.Property) (int, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}
