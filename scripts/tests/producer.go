package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type Message struct {
	UserId      string `json:"userId"`
	MessageBody string `json:"messageBody"`
}

func main() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notification_queue", // queue name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type your messages below. Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		messageBody, _ := reader.ReadString('\n')
		messageBody = strings.TrimSpace(messageBody)

		if messageBody == "exit" {
			fmt.Println("Exiting...")
			break
		}

		message := Message{
			UserId:      "some-user-id", // Replace with actual user ID as needed
			MessageBody: messageBody,
		}

		messageBytes, err := json.Marshal(message)
		if err != nil {
			log.Fatalf("Failed to marshal message to JSON: %v", err)
		}

		err = ch.Publish(
			"",           // exchange
			q.Name,       // routing key (queue name)
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        messageBytes,
			})
		if err != nil {
			log.Fatalf("Failed to publish a message: %v", err)
		}

		log.Printf("Message sent: %s", messageBytes)
	}
}