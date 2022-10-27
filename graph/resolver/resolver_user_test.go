package resolver

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx           context.Context
	userService   *services.UserServices
	logger        *log.Logger
	store         *gorm.DB
	configuration *config.Configuration
	err           error
)

func init() {
	logger := log.New()
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Errorf("panic loading env: %v", err)
	}
	configuration = config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading Test config: %v", err)
	}
	store, _ = database.InitDB()

	userService = services.NewUserService(store, logger, &configuration.JwtConfig)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "config", configuration)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "log", logger)
}

func Test_Resolver_User(t *testing.T) {
	var signIn struct {
		SignIn struct {
			Token string
		}
	}
	var user *model.User

	// get authed test user
	accessToken := makeLoginUser()

	var srv = makeAuthedServer(accessToken, ctx)

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
