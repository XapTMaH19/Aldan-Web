package storage

import (
	"encoding/json"
	"github.com/XapTMaH19/AldanWeb/internal/models"
	"strconv"
)

type AuthRabbitMQ struct {
	client *ClientRabbitMQ
}

type userRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthRabbitMQ(clientRabbitMQ *ClientRabbitMQ) *AuthRabbitMQ {
	clientRabbitMQ.DeclareQueue(clientRabbitMQ.namesQueues.Register)
	clientRabbitMQ.DeclareQueue(clientRabbitMQ.namesQueues.Login)
	return &AuthRabbitMQ{clientRabbitMQ}
}

func (r *AuthRabbitMQ) CreateUser(user models.User) (int, error) {
	response, err := r.client.SendMessage(r.client.namesQueues.Register, user)
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.Atoi(response)

	return userID, nil
}

func (r *AuthRabbitMQ) GetUser(username string, password string) (models.User, error) {

	// Отправляем запрос и получаем ответ
	response, err := r.client.SendMessage(r.client.namesQueues.Login, userRequest{username, password})
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if err := json.Unmarshal([]byte(response), &user); err != nil {
		return models.User{}, err
	}

	return user, nil
}
