package handler

import (
	"encoding/json"
	"net/http"
)

func ResponseStatusOk(w http.ResponseWriter, message string) {
	// Функуия записывает ответ с статус кодом 200 в JSON и возвращает его
	res := make(map[string]string)
	res["message"] = message
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(resJSON)
}

func ResponseError(w http.ResponseWriter, message string, code int) {
	// Функуия записывает ответ с статус кодом 4xx или 5хх в JSON и возвращает его
	res := make(map[string]string)
	res["message"] = message
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(resJSON)
}
