package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQConnection interface {
	ConnectToRabbitMQ() *amqp.Connection
}
type rabbitmq struct {
}

func RabbitMQConnectionInstance() RabbitMQConnection {
	return &rabbitmq{}
}

func (rb *rabbitmq) createRabbitMQConnection() (*amqp.Connection, error) {
	rabbitMQURL := "amqp://username:password@rabbitmq-server:5672/"

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
func (rb *rabbitmq) closeRabbitMQConnection(conn *amqp.Connection) {
	if conn != nil {
		conn.Close()
	}
}

func (rb *rabbitmq) ConnectToRabbitMQ() *amqp.Connection {
	rbmq, err := rb.createRabbitMQConnection()
	if err != nil {
		log.Info("Unable to connect", err)
	}
	return rbmq
}
