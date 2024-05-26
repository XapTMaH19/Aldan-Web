package storage

import (
	"github.com/XapTMaH19/AldanWeb/internal/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string, password string) (models.User, error)
}

type Storage struct {
	Authorization
}

func NewStorage(r *ClientRabbitMQ) *Storage {
	return &Storage{
		Authorization: NewAuthRabbitMQ(r),
	}
}
