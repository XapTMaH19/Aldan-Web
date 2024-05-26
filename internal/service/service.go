package service

import (
	"github.com/XapTMaH19/AldanWeb/internal/models"
	"github.com/XapTMaH19/AldanWeb/internal/storage"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Authorization: NewAuthService(storage.Authorization),
	}
}
