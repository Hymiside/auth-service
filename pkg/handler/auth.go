package handler

import (
	"encoding/json"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/Hymiside/auth-microservice/pkg/service"
	"io/ioutil"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	u := &models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, "Invalid request 1", 400)
		return
	}
	err = json.Unmarshal(body, u)
	if err != nil {
		ResponseError(w, err.Error(), 400)
		return
	}

	msg, err := service.CreateNewUser(u)
	if err != nil {
		ResponseError(w, msg, 400)
	}

	ResponseStatusOk(w, msg)
	return
}
