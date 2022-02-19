package database

import (
	"fmt"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

type ConfigDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// InitDatabase функция инциализирует подключение к базе данных
func InitDatabase(c ConfigDatabase) error {
	var err error

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	db, err = sqlx.Connect("postgres", connect)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return err
}

// CloseDatabase функция закрывает подключение к базе данных
func CloseDatabase() error {
	return db.Close()
}

// ToCreateUser функция добавляет нового пользователя в БД и возвращает ошибку
func ToCreateUser(u models.User) error {
	_, err := db.NamedExec(`INSERT INTO users (uuid, name, username, password_hash) VALUES (:uuid, :name, :username, :password_hash)`, u)
	if err != nil {
		return err
	}
	return err
}

// GetUser функция возвращает uuid и password пользователя
func GetUser(u models.SighInUser) (*sqlx.Rows, error) {
	id, err := db.Queryx(`SELECT uuid, password_hash, username FROM users WHERE username=$1`, u.Username)
	return id, err
}
