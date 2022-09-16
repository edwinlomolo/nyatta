package resolver

import (
	"context"
	"fmt"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.NewUser) (*model.Token, error) {
	token, err := ctx.Value("userService").(*services.UserServices).SignIn(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", nyatta_context.ResolverError, err)
	}
	return &model.Token{Token: *token}, nil
}
