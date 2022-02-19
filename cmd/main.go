package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Hymiside/auth-microservice/pkg/service"

	"github.com/Hymiside/auth-microservice/pkg/config"
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/repository"
	"github.com/Hymiside/auth-microservice/pkg/server"
)

func main() {
	cfgSrv, cfgDb := config.InitConfig()

	srv := &server.Server{}
	h := &handler.Handler{}

	repo, err := repository.NewRepository(cfgDb)
	if err != nil {
		log.Fatalf(err.Error())
	}
	services := service.NewService(*repo)

	go func() {
		if err = srv.RunServer(h.InitHandler(*services), cfgSrv); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	log.Println("authentication microservice launched on http://localhost:5000/")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.ShutdownServer(ctx); err != nil {
		log.Fatalf(err.Error())
	}
	if err = repo.Close(); err != nil {
		log.Fatalf(err.Error())
	}
}
