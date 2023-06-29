package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
)

// Contact is the resolver for the contact field.
func (r *shootResolver) Contact(ctx context.Context, obj *model.Shoot) (*model.Caretaker, error) {
	panic(fmt.Errorf("not implemented: Contact - contact"))
}

// Shoot returns generated.ShootResolver implementation.
func (r *Resolver) Shoot() generated.ShootResolver { return &shootResolver{r} }

type shootResolver struct{ *Resolver }
