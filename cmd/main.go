package main

import (
	"log"
	"net/http"
	"github.com/JabJabHiwHiw/notification-service/internal/notification"
	"github.com/JabJabHiwHiw/notification-service/internal/rabbitmq"
)



func main() {
	// Initialize Firebase client

	// Initialize MongoDB
	err := notification.InitMongoDB("mongodb://root:example@mongodb:27017")
	if err != nil {
		log.Fatalf("failed to init MongoDB: %v\n", err)
	}

	// Start consuming messages from RabbitMQ
	go rabbitmq.ConsumeMessages(func(msg rabbitmq.Message) {
		// Use UserID and Message fields
		err = notification.HandleMessage(msg.UserId, msg.MessageBody)
		if err != nil {
			log.Printf("Error handling message: %v", err)
		}
	})

	// Start HTTP server
	http.HandleFunc("/notifications", notification.GetNotifications)
	log.Println("Notification Service is running on port 8080 with CORS enabled...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}