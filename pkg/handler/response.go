package handler

import (
	"encoding/json"
	"net/http"
)

func responseStatusOk(w http.ResponseWriter) {
	// Функуия записывает ответ с статус кодом 200 в JSON и возвращает его
	res := make(map[string]string)
	res["message"] = "Status ok"
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(resJSON)
}

func responseBadRequest(w http.ResponseWriter) {
	// Функуия записывает ответ с статус кодом 400 в JSON и возвращает его
	res := make(map[string]string)
	res["message"] = "Invalid request"
	resJSON, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	_, _ = w.Write(resJSON)
}
