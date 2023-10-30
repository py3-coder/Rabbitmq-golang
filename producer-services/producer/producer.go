package producer

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func SendProductToQueue(productID int) error {
	err := PublishToQueue(productID)
	if err != nil {
		log.Printf("Failed to send product to queue: %v", err)
	}
	return nil
}

func PublishToQueue(val int) error {
	rabbitMQURL := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Info("Error:", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	queueName := "product_queue"
	_, err = ch.QueueDeclare(
		queueName, 
		false,     
		false,    
		false,     
		false,     
		nil,       
	)
	if err != nil {
		log.Fatal(err)
	}

	message := strconv.Itoa(val)

	// Publish a message to the queue
	err = ch.Publish(
		"",        
		queueName, 
		false,     
		false,     
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sent message: %s", message)
	return nil
}
