package resolver

import (
	"context"
	"fmt"
	"log"
	"testing"

	nyatta_context "github.com/3dw1nM0535/nyatta/context"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ctx           context.Context
	userService   *services.UserServices
	logger        *zap.SugaredLogger
	store         *gorm.DB
	cfg           *nyatta_context.Config
	configuration *nyatta_context.Config
	err           error
)

func init() {
	configuration, err = nyatta_context.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Error reading Test config: %v", err)
	}
	logger, _ = services.NewLogger(configuration)
	store, _ = nyatta_context.OpenDB(configuration, logger)
	userService = services.NewUserService(store, logger, configuration)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "log", logger)
}

func Test_Resolver_User(t *testing.T) {
	var srv = client.New(h.AddContext(ctx, handler.NewDefaultServer(generated.NewExecutableSchema(New()))))
	var signIn struct {
		SignIn struct {
			Token string
		}
	}
	var user *model.User

	t.Run("resolver_should_sign_in_user", func(t *testing.T) {
		query := fmt.Sprintf(`mutation { signIn (input: { first_name: "%s", last_name: "%s", email: "%s" }) { token } }`, "Jane", "Doe", util.GenerateRandomEmail())

		srv.MustPost(query, &signIn)

		assert.NotEqual(t, signIn.SignIn.Token, "")
	})
	t.Run("resolver_should_get_user", func(t *testing.T) {
		var getUser struct {
			GetUser struct{ Email string }
		}

		email := util.GenerateRandomEmail()
		user, _ = userService.CreateUser(&model.NewUser{FirstName: "John", LastName: "Doe", Email: email})
		srv.MustPost(`query ($id: ID!) { getUser (id: $id) { email } }`, &getUser, client.Var("id", user.ID))

		assert.Equal(t, getUser.GetUser.Email, email)
	})
}
