package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// Type is the resolver for the type field.
func (r *propertyResolver) Type(ctx context.Context, obj *model.Property) (model.PropertyType, error) {
	panic(fmt.Errorf("not implemented: Type - type"))
}

// Uploads is the resolver for the uploads field.
func (r *propertyResolver) Uploads(ctx context.Context, obj *model.Property) ([]*model.AnyUpload, error) {
	panic(fmt.Errorf("not implemented: Uploads - uploads"))
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
	panic(fmt.Errorf("not implemented: Caretaker - caretaker"))
}

// Owner is the resolver for the owner field.
func (r *propertyResolver) Owner(ctx context.Context, obj *model.Property) (*model.User, error) {
	return &model.User{}, nil
}

// Property returns generated.PropertyResolver implementation.
func (r *Resolver) Property() generated.PropertyResolver { return &propertyResolver{r} }

type propertyResolver struct{ *Resolver }
