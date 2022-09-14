package resolver

import (
	"context"
	"fmt"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	foundUser, err := ctx.Value("userService").(*services.UserServices).GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", nyatta_context.ResolverError, err)
	}
	return foundUser, nil
}
