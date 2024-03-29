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

// Bedrooms is the resolver for the bedrooms field.
func (r *unitResolver) Bedrooms(ctx context.Context, obj *model.Unit) ([]*model.Bedroom, error) {
	foundBedrooms, err := ctx.Value("unitService").(services.UnitService).GetUnitBedrooms(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return foundBedrooms, nil
}

// Thumbnail is the resolver for the thumbnail field.
func (r *unitResolver) Thumbnail(ctx context.Context, obj *model.Unit) (*model.AnyUpload, error) {
	foundThumbnail, err := ctx.Value("unitService").(services.UnitService).GetUnitThumbnail(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return foundThumbnail, nil
}

// Caretaker is the resolver for the caretaker field.
func (r *unitResolver) Caretaker(ctx context.Context, obj *model.Unit) (*model.Caretaker, error) {
	if obj.CaretakerID == nil {
		return nil, nil
	}
	caretaker, err := ctx.Value("propertyService").(services.PropertyService).GetPropertyCaretaker(ctx, *obj.CaretakerID)
	if err != nil {
		return nil, err
	}

	return caretaker, nil
}

// Property is the resolver for the property field.
func (r *unitResolver) Property(ctx context.Context, obj *model.Unit) (*model.Property, error) {
	property, err := ctx.Value("propertyService").(services.PropertyService).GetProperty(ctx, obj.PropertyID)
	if err != nil {
		return nil, err
	}

	return property, nil
}

// Tenant is the resolver for the tenant field.
func (r *unitResolver) Tenant(ctx context.Context, obj *model.Unit) (*model.Tenant, error) {
	tenant, err := ctx.Value("tenancyService").(services.TenancyService).GetCurrentTenant(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return tenant, err
}

// Owner is the resolver for the owner field.
func (r *unitResolver) Owner(ctx context.Context, obj *model.Unit) (*model.User, error) {
	if obj.CreatedBy == nil {
		return nil, nil
	}

	user, err := ctx.Value("userService").(services.UserService).GetUser(ctx, *obj.CreatedBy)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Amenities is the resolver for the amenities field.
func (r *unitResolver) Amenities(ctx context.Context, obj *model.Unit) ([]*model.Amenity, error) {
	amenities, err := ctx.Value("amenityService").(services.AmenityService).GetUnitAmenities(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return amenities, err
}

// Images is the resolver for the images field.
func (r *unitResolver) Images(ctx context.Context, obj *model.Unit) ([]*model.AnyUpload, error) {
	uploads, err := ctx.Value("unitService").(services.UnitService).GetUnitImages(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return uploads, nil
}

// Tenancy is the resolver for the tenancy field.
func (r *unitResolver) Tenancy(ctx context.Context, obj *model.Unit) ([]*model.Tenant, error) {
	foundTenancies, err := ctx.Value("tenancyService").(services.TenancyService).GetUnitTenancy(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return foundTenancies, nil
}

// Unit returns generated.UnitResolver implementation.
func (r *Resolver) Unit() generated.UnitResolver { return &unitResolver{r} }

type unitResolver struct{ *Resolver }
