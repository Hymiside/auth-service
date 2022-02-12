package main

import (
	"github.com/Hymiside/auth-microservice/pkg/database"
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"log"
)

func main() {
	srv := new(server.Server)
	handlers := new(handler.Handler)

	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			log.Fatalf("Error closing database connection: %s", err.Error())
		}
	}()

	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Database initialization error: %s", err.Error())
	}

	if err = srv.RunServer(handlers.InitHandler()); err != nil {
		log.Fatalf("Server crashed on startup: %s", err.Error())
	}
}
