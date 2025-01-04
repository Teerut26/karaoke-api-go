package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func Firestore() *firestore.Client {
	app := Firebase()
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Firestore client: %v", err)
	}
	return client
}
