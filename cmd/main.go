package main

import (
    "log"
	"context"

    "github.com/JabJabHiwHiw/notification-service/config"
    "github.com/JabJabHiwHiw/notification-service/internal/rabbitmq"
    "github.com/JabJabHiwHiw/notification-service/internal/firebase"
    "github.com/JabJabHiwHiw/notification-service/internal/notification"
)

func main() {
    config.LoadConfig()

	client, err := event.InitFirebaseClient(context.Background())
	if err != nil {
		log.Fatalf("failed to init firebase client: %v\n", err)
	}

    rabbitmq.ConsumeMessages(func(msg string) {
        err := notification.HandleMessage(client, msg)
        if err != nil {
            log.Printf("Error handling message: %v", err)
        }
    })
}