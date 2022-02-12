package service

import (
	"github.com/Hymiside/auth-microservice/pkg/database"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword функция хэширует пароль, а затем возвращает пароль и ошибку
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash функция хэширует пароль и возравщает хэш пароля из БД, а затем возвращает булевое значение
func CheckPasswordHash(passwordHash, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordHash))
	return err == nil
}

// CreateNewUser хэширует пароль, меняет его в структуре User и передает её в виде аргумента
func CreateNewUser(u *models.User) (string, error) {
	passwordHash, err := HashPassword(u.Password)
	if err != nil {
		return "Internal server error", err
	}

	u.Password = passwordHash
	err = database.ToCreateUser(u)
	if err != nil {
		return "Internal server error", err
	}

	return "User successfully sign up", nil
}
