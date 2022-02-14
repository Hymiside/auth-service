package handler

import (
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler *chi.Mux
}

// InitHandler функция инициализирует обработчики
func (h *Handler) InitHandler() *chi.Mux {
	h.handler = chi.NewRouter()

	h.handler.Post("/api/auth/sign-up", signUp)
	h.handler.Post("/api/auth/sign-in", signIn)

	return h.handler
}
