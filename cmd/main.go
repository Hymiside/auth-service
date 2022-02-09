package main

import (
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"log"
)

import (
	"github.com/Hymiside/auth-microservice/pkg/database"
	"github.com/Hymiside/auth-microservice/pkg/handler"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"github.com/Hymiside/auth-microservice/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инциализации конфига: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	handlers := new(handler.Handler)

	if err := srv.RunServer(handlers.InitHandler()); err != nil {
		log.Fatalf("Сервер упал при запуске: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
