package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ConfigServer struct {
	Host string
	Port string
}

type Server struct {
	httpServer *http.Server
}

// RunServer запускает HTTP сервер
func (s *Server) RunServer(handler *chi.Mux, c ConfigServer) error {

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

// ShutdownServer выключает HTTP сервер
func (s *Server) ShutdownServer(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	return err
}
