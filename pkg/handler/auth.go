package handler

import (
	"encoding/json"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"io/ioutil"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	m := &models.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseBadRequest(w)
		return
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		responseBadRequest(w)
		return
	}
	if m.Username == nil || m.Name == nil || m.Password == nil {
		responseBadRequest(w)
		return
	}
	// TODO

	responseStatusOk(w)
	return
}
