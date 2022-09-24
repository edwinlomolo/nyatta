package resolver

import (
	"context"
	"fmt"
	"testing"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
)

func Test_unauthed_graphql_request(t *testing.T) {
	var signIn struct {
		SignIn struct {
			Token string
		}
	}
	var srv = client.New(h.AddContext(context.Background(), h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))))

	t.Run("should_not_next_unauthed_graphql_request", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { signIn (input: { first_name: "%s", last_name: "%s", email: "%s" }) { token } }`, "Jane", "Doe", util.GenerateRandomEmail())

		err := srv.Post(query, &signIn)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "http 401: Unauthorized")
	})
}
