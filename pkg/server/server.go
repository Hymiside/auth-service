package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "5000"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(handler *chi.Mux) error {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: handler,
	}
	log.Println("The auth microservice is running on http://localhost:5000/")
	return s.httpServer.ListenAndServe()
}
