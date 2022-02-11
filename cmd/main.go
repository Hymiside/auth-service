package main

import (
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"log"
)

func main() {
	srv := new(server.Server)
	handlers := new(handler.Handler)

	if err := srv.RunServer(handlers.InitHandler()); err != nil {
		log.Fatalf("Server crashed on startup: %s", err.Error())
	}
}
