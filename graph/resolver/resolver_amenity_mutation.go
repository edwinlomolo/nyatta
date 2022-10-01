package resolver

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/model"
)

// AddAmenity is the resolver for the addAmenity field.
func (r *mutationResolver) AddAmenity(ctx context.Context, input model.AmenityInput) (*model.Amenity, error) {
	panic(fmt.Errorf("not implemented: AddAmenity - addAmenity"))
}
