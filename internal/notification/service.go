package notification

import (
    "context"
    "firebase.google.com/go/messaging"
    "log"
)

func HandleMessage(client *messaging.Client, messageBody string) error {
    message := &messaging.Message{
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