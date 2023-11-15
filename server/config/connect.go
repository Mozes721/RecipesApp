package config

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
)

func FirebaseApp(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile("/mnt/c/Users/RichardTaujenis/Desktop/RecipesApp/server/config/account_key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	app, err := FirebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAuthClient(ctx context.Context) (*auth.Client, error) {
	app, err := FirebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return authClient, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
