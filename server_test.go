package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	h "github.com/3dw1nM0535/nyatta/handler"
	"github.com/stretchr/testify/assert"
)

func Test_root_api(t *testing.T) {
	reqData := "Hello, world"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, reqData)
	})

	t.Run("should_drop_unauthed_api_request", func(t *testing.T) {
		// TODO rfr to makeAuthedServer
		srv := httptest.NewServer(h.Authenticate(handler))
		defer srv.Close()

		req, err := http.NewRequest(http.MethodPost, srv.URL, nil)
		assert.Nil(t, err)

		c := srv.Client()
		res, err := c.Do(req)
		assert.Nil(t, err)
		defer res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusUnauthorized)

		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, string(data), "Unauthorized\n")
	})
}
