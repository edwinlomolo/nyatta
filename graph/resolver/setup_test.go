package resolver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	log "github.com/sirupsen/logrus"
)

// makeLoginUser - return authed user
func makeLoginUser() string {
	var creds struct {
		AccessToken string `json:"access_token"`
		Code        int    `json:"code"`
	}
	var jsonStr = []byte(fmt.Sprintf(`{"first_name": %q, "last_name": %q, "email": %q, "avatar": %q}`, "john", "doe", util.GenerateRandomEmail(), "https://avatar.jpg"))

	httpServer := httptest.NewServer(h.AddContext(ctx, h.Handshake()))
	defer httpServer.Close()

	url := fmt.Sprintf("%s/handshake", httpServer.URL)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))

	c := httpServer.Client()
	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(data, &creds)
	return creds.AccessToken
}

// makeAuthedGqlServer - return authenticated graphql client
func makeAuthedGqlServer(authenticate bool, ctx context.Context) *client.Client {
	var srv *client.Client
	if !authenticate {
		// unauthed client
		srv = client.New(h.AddContext(ctx, handler.NewDefaultServer(generated.NewExecutableSchema(New()))))
		return srv
	}
	// authed user
	tokenString := makeLoginUser()
	// authed client
	srv = client.New(h.AddContext(ctx, h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))), client.AddHeader("Authorization", fmt.Sprintf("Bearer %s", tokenString)))

	return srv
}
