package service

import "github.com/Hymiside/auth-microservice/pkg/database"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
