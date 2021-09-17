package config

import (
	"context"

	firestore "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Connection() (*firestore.App, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("secret.json")
	return firestore.NewApp(ctx, nil, opt)
}
