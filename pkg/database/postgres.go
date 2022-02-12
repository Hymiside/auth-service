package database

import (
	"fmt"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "putinbest"
	dbname   = "auth_microservice_db"
)

// InitDatabase функция инциализирует подключение к базе данных
func InitDatabase() error {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", connect)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// CloseDatabase фунция закрывает подключение к базе данных
func CloseDatabase() error {
	return db.Close()
}

// ToCreateUser Функция дабвляет нового пользователя в БД и возвращает ошибку
func ToCreateUser(u *models.User) error {
	_, err := db.NamedExec(`INSERT INTO users (name, username, password_hash) VALUES (:name, :username, :password_hash)`, u)
	if err != nil {
		return err
	}
	return nil
}
