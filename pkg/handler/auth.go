package handler

import (
	"encoding/json"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/Hymiside/auth-microservice/pkg/service"
	"net/http"
)

// signUp функция регистрации нового пользователя
// парсит запрос, проверяет поля и отправляет структуру в виду аргумента
func signUp(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ResponseError(w, "ошибка декордирования запроса", 400)
		return
	}
	if u.Name == "" || u.Username == "" || u.Password == "" {
		ResponseError(w, "некорректный формат запроса", 400)
		return
	}

	msg, err := service.CreateNewUser(u)
	if err != nil {
		ResponseError(w, msg, 400)
		return
	}
	ResponseStatusOk(w, "uuid", msg)
	return
}

// signIn функция входа существующего пользователя
// парсит запрос, проверяет поля, отправляет структуру в виду аргумента и возращает либо ошибку либо токен
func signIn(w http.ResponseWriter, r *http.Request) {
	var u models.SighInUser

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		ResponseError(w, "ошибка декордирования запроса", 400)
		return
	}
	if u.Username == "" || u.Password == "" {
		ResponseError(w, "некорректный формат запроса", 400)
		return
	}

	msg, err := service.CheckUser(u)
	if err != nil {
		ResponseError(w, msg, 400)
		return
	}
	ResponseStatusOk(w, "token", msg)
	return
}
