package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
)

// Category is the resolver for the category field.
func (r *amenityResolver) Category(ctx context.Context, obj *model.Amenity) (*string, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Amenity returns generated.AmenityResolver implementation.
func (r *Resolver) Amenity() generated.AmenityResolver { return &amenityResolver{r} }

type amenityResolver struct{ *Resolver }
