package config

import (
	"context"
	"encoding/json"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func Firebase() *firebase.App {
	configJSON := os.Getenv("FIREBASE_ADMIN")
	if configJSON == "" {
		log.Fatal("FIREBASE_ADMIN environment variable is not set")
	}

	var config FirebaseConfig
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		log.Fatalf("Failed to parse Firebase configuration: %v", err)
	}

	// Convert the config to a JSON byte array
	configBytes, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("Failed to marshal Firebase configuration: %v", err)
	}

	opt := option.WithCredentialsJSON(configBytes)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Admin SDK: %v", err)
	}

	return app
}
