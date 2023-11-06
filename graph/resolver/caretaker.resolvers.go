package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
)

// Uploads is the resolver for the uploads field.
func (r *caretakerResolver) Uploads(ctx context.Context, obj *model.Caretaker) ([]*model.AnyUpload, error) {
	panic(fmt.Errorf("not implemented: Uploads - uploads"))
}

// Properties is the resolver for the properties field.
func (r *caretakerResolver) Properties(ctx context.Context, obj *model.Caretaker) ([]*model.Property, error) {
	panic(fmt.Errorf("not implemented: Properties - properties"))
}

// Caretaker returns generated.CaretakerResolver implementation.
func (r *Resolver) Caretaker() generated.CaretakerResolver { return &caretakerResolver{r} }

type caretakerResolver struct{ *Resolver }
