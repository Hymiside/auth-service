package handler

import (
	"github.com/Hymiside/auth-microservice/pkg/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler *chi.Mux
}

// InitHandler функция инициализирует обработчики
func (h *Handler) InitHandler(s service.Service) *chi.Mux {
	h.handler = chi.NewRouter()
	handlers := NewHandlers(s)

	h.handler.Post("/api/auth/sign-up", handlers.signUp)
	h.handler.Post("/api/auth/sign-in", handlers.signIn)

	return h.handler
}
