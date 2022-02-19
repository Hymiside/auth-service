package repository

import (
	"fmt"

	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Repository struct {
	db *sqlx.DB
}

// NewRepository функция инциализирует подключение к базе данных
func NewRepository(c ConfigDatabase) (*Repository, error) {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	db, err := sqlx.Connect("postgres", connect)
	if err != nil {
		return nil, fmt.Errorf("failed to create db conn: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connection test error: %w", err)
	}
	return &Repository{db: db}, err
}

// Close функция закрывает подключение к базе данных
func (r *Repository) Close() error {
	err := r.db.Close()
	if err != nil {
		return fmt.Errorf("error closing database connection: %w", err)
	}
	return nil
}

// ToCreateUser функция добавляет нового пользователя в БД и возвращает ошибку
func (r *Repository) ToCreateUser(u models.User) error {
	_, err := r.db.NamedExec(`INSERT INTO users (uuid, name, username, password_hash) VALUES (:uuid, :name, :username, :password_hash)`, u)
	if err != nil {
		return fmt.Errorf("new user creation error: %w", err)
	}
	return nil
}

// GetUser функция возвращает uuid и password пользователя
func (r *Repository) GetUser(username string) (models.User, error) {
	var u models.User

	row := r.db.QueryRow(`SELECT uuid, password_hash, username FROM users WHERE username = $1`, username)
	if err := row.Scan(&u.UUID, &u.Password, &u.Username); err != nil {
		return models.User{}, fmt.Errorf("failed to scan row: %w", err)
	}

	if err := row.Err(); err != nil {
		return models.User{}, fmt.Errorf("got rows error: %w", err)
	}
	return u, nil
}
