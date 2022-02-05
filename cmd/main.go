package main

import (
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"log"
)

func main() {
	h := new(handler.Handler)
	srv := new(server.Server)
	err := srv.RunServer(h.InitRoutes())

	if err != nil {
		log.Fatal("Сервер упал при запуске", err)
	}
}
