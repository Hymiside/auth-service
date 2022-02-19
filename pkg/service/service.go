package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Hymiside/auth-microservice/pkg/models"
	"github.com/Hymiside/auth-microservice/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	signingKey = []byte("qrkjk#4#%35FSFJlja#4353KSFjH")
)

type Service struct {
	repo *repository.Repository
}

func NewService(r repository.Repository) *Service {
	return &Service{repo: &r}
}

// CreateNewUser хэширует пароль, меняет его в структуре User и передает её в виде аргумента
func (s *Service) CreateNewUser(u models.User) (string, error) {
	var err error

	u.Password, err = hashPassword(u.Password)
	if err != nil {
		return "", fmt.Errorf("password hash error: %w", err)
	}
	u.UUID = uuid.New().String()

	err = s.repo.ToCreateUser(u)
	if err != nil {
		return "", err
	}

	return u.UUID, nil
}

// CheckUser функция проверяет выполняет проверку для входа пользователя
func (s *Service) CheckUser(u models.SighInUser) (string, error) {
	user, err := s.repo.GetUser(u.Username)
	if err != nil {
		return "", err
	}

	userAuth := checkPasswordHash(u.Password, user.Password)
	if !userAuth || u.Username != user.Username {
		return "", errors.New("user entered an incorrect password")
	}

	timeNow := time.Now()

	claims := models.TokenClaims{
		UserId: user.UUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeNow.Add(12 * time.Hour).Unix(),
			IssuedAt:  timeNow.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signingKey)
}

// hashPassword функция хэширует пароль, а затем возвращает пароль и ошибку
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash функция хэширует пароль и сравнивает хэш пароля из БД, а затем возвращает булевое значение
func checkPasswordHash(passwordHash, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordHash)) == nil
}
