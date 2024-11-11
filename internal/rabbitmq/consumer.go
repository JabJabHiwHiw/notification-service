package rabbitmq

import (
	"log"
	"encoding/json"

	"github.com/streadway/amqp"
)

type Message struct {
	UserId      string `json:"userId"`
	MessageBody string `json:"messageBody"`
}

func ConsumeMessages(callback func(msg Message)) {
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"notification_queue", // queue
		"",                   // consumer tag
		true,                 // auto-acknowledge
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			var message Message
			// Unmarshal the JSON message into the Message struct
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			// Pass the Message struct to the callback
			callback(message)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	select {}
}