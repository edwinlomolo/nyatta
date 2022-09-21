package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

func Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		loginResponse := &model.LoginResponse{}
		var newUser *model.NewUser

		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &newUser)
		token, err := ctx.Value("userService").(*services.UserServices).SignIn(newUser)
		if err != nil {
			response := &model.Response{
				Code: http.StatusInternalServerError,
				Err:  err.Error(),
			}
			loginResponse.Response = response
			writeResponse(w, loginResponse, loginResponse.Code)
			return
		}

		response := &model.Response{
			Code: http.StatusOK,
		}
		loginResponse.Response = response
		loginResponse.AccessToken = *token

		writeResponse(w, loginResponse, loginResponse.Code)
	})
}

func writeResponse(w http.ResponseWriter, response interface{}, code int) {
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
