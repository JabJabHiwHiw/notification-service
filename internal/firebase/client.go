package event

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func InitFirebaseClient(ctx context.Context) (*messaging.Client, error) {
	filePath := "./internal/firebase/secrets/sa-hiw-hiw-2b446b26808d.json"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("error file not found: %v\n", err)
	}

	opt := option.WithCredentialsFile(filePath)
	fmt.Println("opt", opt)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting messaging client: %v\n", err)
		return nil, err
	}

	return client, nil
}
