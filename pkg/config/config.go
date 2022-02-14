package config

import (
	"github.com/Hymiside/auth-microservice/pkg/database"
	"github.com/Hymiside/auth-microservice/pkg/server"
	"github.com/joho/godotenv"
	"os"
)

func InitConfig() (server.ConfigServer, database.ConfigDatabase) {
	_ = godotenv.Load()

	host, _ := os.LookupEnv("SERVICE_HOST")
	port, _ := os.LookupEnv("SERVICE_PORT")

	hostDb, _ := os.LookupEnv("DB_HOST")
	portDb, _ := os.LookupEnv("DB_PORT")
	userDb, _ := os.LookupEnv("DB_USER")
	passwordDb, _ := os.LookupEnv("DB_PASSWORD")
	nameDb, _ := os.LookupEnv("DB_NAME")

	configDb := database.ConfigDatabase{
		Host:     hostDb,
		Port:     portDb,
		User:     userDb,
		Password: passwordDb,
		Name:     nameDb,
	}

	config := server.ConfigServer{
		Host: host,
		Port: port,
	}
	return config, configDb
}
