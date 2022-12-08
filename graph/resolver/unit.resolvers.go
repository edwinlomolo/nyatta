package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
)

// Bedrooms is the resolver for the bedrooms field.
func (r *propertyUnitResolver) Bedrooms(ctx context.Context, obj *model.PropertyUnit) ([]*model.Bedroom, error) {
	panic(fmt.Errorf("not implemented: Bedrooms - bedrooms"))
}

// Tenancy is the resolver for the tenancy field.
func (r *propertyUnitResolver) Tenancy(ctx context.Context, obj *model.PropertyUnit) ([]*model.Tenant, error) {
	panic(fmt.Errorf("not implemented: Tenancy - tenancy"))
}

// PropertyUnit returns generated.PropertyUnitResolver implementation.
func (r *Resolver) PropertyUnit() generated.PropertyUnitResolver { return &propertyUnitResolver{r} }

type propertyUnitResolver struct{ *Resolver }
