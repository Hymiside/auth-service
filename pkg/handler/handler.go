package handler

import (
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler *chi.Mux
}

func (h *Handler) InitHandler() *chi.Mux {
	h.handler = chi.NewRouter()

	h.handler.Post("/api/auth/signup", signUp)

	return h.handler
}
