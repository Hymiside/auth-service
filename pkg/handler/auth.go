package handler

import (
	"encoding/json"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/Hymiside/auth-microservice/pkg/service"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ResponseError(w, "Invalid request 1", 400)
		return
	}

	msg, err := service.CreateNewUser(u)
	if err != nil {
		ResponseError(w, msg, 400)
	}

	ResponseStatusOk(w, msg)
	return
}
