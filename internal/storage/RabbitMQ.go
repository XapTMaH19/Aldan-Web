package storage

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Host              string
	Port              string
	Username          string
	Password          string
	RegisterNameQueue string
	LoginNameQueue    string
}

type RabbitMQQueues struct {
	Register string `yaml:"register"`
	Login    string `yaml:"login"`
}

type ClientRabbitMQ struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	namesQueues *RabbitMQQueues
}

func NewClientRabbitMQ(cfg RabbitMQConfig) (*ClientRabbitMQ, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		logrus.Fatalf("произошла ошибка подключения к Rabbit-у: %v", err.Error())
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	r := &RabbitMQQueues{cfg.RegisterNameQueue, cfg.LoginNameQueue}
	return &ClientRabbitMQ{conn: conn, channel: ch, namesQueues: r}, nil
}

func (client *ClientRabbitMQ) DeclareQueue(queueName string) (amqp.Queue, error) {
	return client.channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}

func (client *ClientRabbitMQ) SendMessage(queueName string, message interface{}) (string, error) {
	correlationID := uuid.New().String()

	body, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	err = client.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          body,
			CorrelationId: correlationID,
			ReplyTo:       queueName,
		})
	if err != nil {
		return "", err
	}

	msgs, err := client.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return "", err
	}

	for msg := range msgs {
		if msg.CorrelationId == correlationID {
			return string(msg.Body), nil
		}
	}

	return "", nil
}
