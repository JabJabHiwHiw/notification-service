package notification

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
)

func HandleMessage(client *messaging.Client, messageBody string) error {
	message := &messaging.Message{
		Topic: "notification",
		Notification: &messaging.Notification{
			Title: "New Notification",
			Body:  messageBody,
		},
	}

	_, err := client.Send(context.Background(), message)
	if err != nil {
		log.Printf("Error sending push notification: %v", err)
	} else {
		log.Printf("Push notification sent successfully!")
	}
	return err
}
