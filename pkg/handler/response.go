package handler

import (
	"encoding/json"
	"net/http"
)

func responseStatusOk(w http.ResponseWriter, message string) {
	// Функуия записывает ответ с статус кодом 200 в JSON и возвращает его
	res := make(map[string]string)
	res["message"] = message
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(resJSON)
}

func responseError(w http.ResponseWriter, message string, code int) {
	// Функуия записывает ответ с статус кодом 4xx или 5хх в JSON и возвращает его
	res := make(map[string]string)
	res["message"] = message
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(resJSON)
}
