package resolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func makeLoginUser() string {
	var creds struct {
		AccessToken string `json:"access_token"`
		Code        int    `json:"code"`
	}
	var jsonStr = []byte(fmt.Sprintf(`{"first_name": "%s", "last_name": "%s", "email": "%s"}`, "john", "doe", util.GenerateRandomEmail()))

	httpServer := httptest.NewServer(h.AddContext(ctx, h.Login()))
	defer httpServer.Close()

	url := fmt.Sprintf("%s/login", httpServer.URL)
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

func makeAuthedServer(tokenString string) *client.Client {
	var srv *client.Client
	if len(tokenString) == 0 {
		// return unauthed client
		srv = client.New(h.AddContext(ctx, h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))))
	} else {
		// return authed client
		srv = client.New(h.AddContext(ctx, h.Authenticate(handler.NewDefaultServer(generated.NewExecutableSchema(New())))), client.AddHeader("Authorization", fmt.Sprintf("Bearer %s", tokenString)))
	}
	return srv
}
