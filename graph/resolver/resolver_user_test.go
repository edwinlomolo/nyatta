package resolver

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/database"
	sqlStore "github.com/3dw1nM0535/nyatta/database/store"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var (
	ctx             context.Context
	userService     *services.UserServices
	propertyService *services.PropertyServices
	logger          *log.Logger
	configuration   *config.Configuration
	err             error
	db              *sql.DB
)

func TestMain(m *testing.M) {
	// setup tests
	logger := log.New()
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Errorf("panic loading env: %v", err)
	}
	configuration = config.LoadConfig()
	db, err = database.InitDB()
	if err != nil {
		log.Fatalf("%s: %v", database.DatabaseError, err)
	}
	queries := sqlStore.New(db)

	userService = services.NewUserService(queries, logger, &configuration.JwtConfig)
	propertyService = services.NewPropertyService(queries, logger)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "config", configuration)
	ctx = context.WithValue(ctx, "userService", userService)
	ctx = context.WithValue(ctx, "propertyService", propertyService)
	ctx = context.WithValue(ctx, "log", logger)

	// exit once done
	os.Exit(m.Run())
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
			GetUser struct {
				Email      string
				First_Name string
			}
		}
		var err error

		email := util.GenerateRandomEmail()
		user, err = userService.CreateUser(&model.NewUser{FirstName: "John", LastName: "Doe", Email: email})
		if err != nil {
			t.Errorf("expected nil err got %v", err)
		}
		srv.MustPost(`query ($id: ID!) { getUser (id: $id) { email, first_name } }`, &getUser, client.Var("id", user.ID))

		log.Infoln(getUser.GetUser)
		assert.Equal(t, getUser.GetUser.Email, email)
		assert.Equal(t, getUser.GetUser.First_Name, "John")
	})
}
