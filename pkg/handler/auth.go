package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/Hymiside/auth-microservice/pkg/service"
)

type Handlers struct {
	serv *service.Service
}

func NewHandlers(s service.Service) *Handlers {
	return &Handlers{serv: &s}
}

// signUp функция регистрации нового пользователя
// парсит запрос, проверяет поля и отправляет структуру в виду аргумента
func (s *Handlers) signUp(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ResponseError(w, fmt.Sprintf("request decorating error: %s", err), 400)
		return
	}
	if u.Name == "" || u.Username == "" || u.Password == "" {
		ResponseError(w, "invalid request format", 400)
		return
	}

	uuid, err := s.serv.CreateNewUser(u)
	if err != nil {
		ResponseError(w, err.Error(), 400)
		return
	}
	ResponseStatusOk(w, "uuid", uuid)
	return
}

// signIn функция входа существующего пользователя
// парсит запрос, проверяет поля, отправляет структуру в виду аргумента и возращает либо ошибку либо токен
func (s *Handlers) signIn(w http.ResponseWriter, r *http.Request) {
	var u models.SighInUser

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ResponseError(w, fmt.Sprintf("request decorating error: %s", err), 400)
		return
	}
	if u.Username == "" || u.Password == "" {
		ResponseError(w, "invalid request format", 400)
		return
	}

	token, err := s.serv.CheckUser(u)
	if err != nil {
		ResponseError(w, err.Error(), 400)
		return
	}
	ResponseStatusOk(w, "token", token)
	return
}
