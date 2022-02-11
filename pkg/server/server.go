package server

import (
	"github.com/go-chi/chi/v5"
	"log"
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
	log.Println("The auth microservice is running on http://localhost:5000/")
	return s.httpServer.ListenAndServe()
}
