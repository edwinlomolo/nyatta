package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// Bedrooms is the resolver for the bedrooms field.
func (r *propertyUnitResolver) Bedrooms(ctx context.Context, obj *model.PropertyUnit) ([]*model.Bedroom, error) {
	foundBedrooms, err := ctx.Value("unitService").(*services.UnitServices).GetUnitBedrooms(obj.ID)
	if err != nil {
		return nil, err
	}
	return foundBedrooms, nil
}

// AmenityCount is the resolver for the amenityCount field.
func (r *propertyUnitResolver) AmenityCount(ctx context.Context, obj *model.PropertyUnit) (int, error) {
	totalAmenities, err := ctx.Value("unitService").(*services.UnitServices).AmenityCount(obj.ID)
	if err != nil {
		return 0, err
	}
	return int(totalAmenities), nil
}

// Uploads is the resolver for the uploads field.
func (r *propertyUnitResolver) Uploads(ctx context.Context, obj *model.PropertyUnit) ([]*model.AnyUpload, error) {
	id, err := strconv.ParseInt(obj.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	uploads, err := ctx.Value("unitService").(*services.UnitServices).GetUnitImages(id)
	if err != nil {
		return nil, err
	}
	return uploads, nil
}

// Tenancy is the resolver for the tenancy field.
func (r *propertyUnitResolver) Tenancy(ctx context.Context, obj *model.PropertyUnit) ([]*model.Tenant, error) {
	foundTenancies, err := ctx.Value("unitService").(*services.UnitServices).GetUnitTenancy(obj.ID)
	if err != nil {
		return nil, err
	}
	return foundTenancies, nil
}

// PropertyUnit returns generated.PropertyUnitResolver implementation.
func (r *Resolver) PropertyUnit() generated.PropertyUnitResolver { return &propertyUnitResolver{r} }

type propertyUnitResolver struct{ *Resolver }
