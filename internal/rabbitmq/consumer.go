package rabbitmq

import (
    "log"
    "github.com/streadway/amqp"
)

func ConsumeMessages(callback func(msg string)) {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
            callback(string(d.Body))
        }
    }()

    log.Printf("Waiting for messages. To exit press CTRL+C")
    select {}
}