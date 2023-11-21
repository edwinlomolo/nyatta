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

// Thumbnail is the resolver for the thumbnail field.
func (r *propertyResolver) Thumbnail(ctx context.Context, obj *model.Property) (*model.AnyUpload, error) {
	thumbnail, err := ctx.Value("propertyService").(*services.PropertyServices).GetPropertyThumbnail(obj.ID)
	if err != nil {
		return nil, err
	}
	return thumbnail, nil
}

// Units is the resolver for the units field.
func (r *propertyResolver) Units(ctx context.Context, obj *model.Property) ([]*model.PropertyUnit, error) {
	foundUnits, err := ctx.Value("propertyService").(*services.PropertyServices).GetPropertyUnits(obj.ID)
	if err != nil {
		return nil, err
	}
	return foundUnits, nil
}

// Caretaker is the resolver for the caretaker field.
func (r *propertyResolver) Caretaker(ctx context.Context, obj *model.Property) (*model.Caretaker, error) {
	caretaker, err := ctx.Value("propertyService").(*services.PropertyServices).GetPropertyCaretaker(obj.CaretakerID)
	if err != nil {
		return nil, err
	}

	return caretaker, nil
}

// Owner is the resolver for the owner field.
func (r *propertyResolver) Owner(ctx context.Context, obj *model.Property) (*model.User, error) {
	return &model.User{}, nil
}

// Property returns generated.PropertyResolver implementation.
func (r *Resolver) Property() generated.PropertyResolver { return &propertyResolver{r} }

type propertyResolver struct{ *Resolver }
