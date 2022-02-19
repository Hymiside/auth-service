package service

import (
	"errors"
	"github.com/Hymiside/auth-microservice/pkg/database"
	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	signingKey                 = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")
	id, passwordHash, username string
)

// hashPassword функция хэширует пароль, а затем возвращает пароль и ошибку
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash функция хэширует пароль и сравнивает хэш пароля из БД, а затем возвращает булевое значение
func checkPasswordHash(passwordHash, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordHash))
	return err == nil
}

// CreateNewUser хэширует пароль, меняет его в структуре User и передает её в виде аргумента
func CreateNewUser(u models.User) (string, error) {
	var err error

	u.Password, err = hashPassword(u.Password)
	if err != nil {
		return err.Error(), err
	}
	u.UUID = uuid.New().String()[:8]

	err = database.ToCreateUser(u)
	if err != nil {
		return err.Error(), err
	}

	return u.UUID, err
}

// CheckUser функция проверяет выполняет проверку для входа пользователя
func CheckUser(u models.SighInUser) (string, error) {
	row, err := database.GetUser(u)
	if err != nil {
		return err.Error(), err
	}

	for row.Next() {
		err = row.Scan(&id, &passwordHash, &username)
		if err != nil {
			return err.Error(), err
		}
	}

	userAuth := checkPasswordHash(u.Password, passwordHash)
	if !userAuth || u.Username != username {
		err = errors.New("пользователь ввел неверный пароль или логин")
		return err.Error(), err
	}

	claims := models.TokenClaims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signingKey)
}
