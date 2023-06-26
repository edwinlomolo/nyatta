package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
)

// ShootsInCharge is the resolver for the shootsInCharge field.
func (r *caretakerResolver) ShootsInCharge(ctx context.Context, obj *model.Caretaker) ([]*model.Shoot, error) {
	panic(fmt.Errorf("not implemented: ShootsInCharge - shootsInCharge"))
}

// Caretaker returns generated.CaretakerResolver implementation.
func (r *Resolver) Caretaker() generated.CaretakerResolver { return &caretakerResolver{r} }

type caretakerResolver struct{ *Resolver }
