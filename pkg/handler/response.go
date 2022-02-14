package handler

import (
	"encoding/json"
	"net/http"
)

// ResponseStatusOk функция возращает ответ со статус кодом 200
func ResponseStatusOk(w http.ResponseWriter, field, message string) {
	res := make(map[string]string)
	res[field] = message
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(resJSON)
}

// ResponseError функция возвращает ответ со статус кодом 4хх или 5хх
func ResponseError(w http.ResponseWriter, message string, code int) {
	res := make(map[string]string)
	res["message"] = message
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(resJSON)
}
