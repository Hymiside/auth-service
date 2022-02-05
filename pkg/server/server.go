package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":5000",
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) CloseServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
