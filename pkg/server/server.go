package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(handler *chi.Mux) error {
	s.httpServer = &http.Server{
		Addr:    ":5000",
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}
