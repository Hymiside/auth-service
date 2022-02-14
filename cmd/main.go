package main

import (
	"context"
	"github.com/Hymiside/auth-microservice/pkg/config"
	"github.com/Hymiside/auth-microservice/pkg/database"
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfgSrv, cfgDb := config.InitConfig()

	srv := new(server.Server)
	handlers := new(handler.Handler)

	err := database.InitDatabase(cfgDb)
	if err != nil {
		log.Fatalf("Ошибка инициализации подключенния к базе данных: %s", err.Error())
	}

	go func() {
		if err = srv.RunServer(handlers.InitHandler(), cfgSrv); err != nil {
			log.Fatalf("Микросервис упал: %s", err.Error())
		}
	}()
	log.Println("Микросервис аутентификации запущен http://localhost:5000/ .")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.ShutdownServer(ctx); err != nil {
		log.Fatalf("Ошибка при остановки микросервиса: %s", err.Error())
	}
	if err = database.CloseDatabase(); err != nil {
		log.Fatalf("Ошибка закрытия подключения к базе данных: %s", err.Error())
	}
	log.Println("Микросервис успешно остановлен.")
}
