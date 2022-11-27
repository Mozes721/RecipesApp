package config

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
)

func FirebaseDB(ctx context.Context) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("/home/mozes/serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, _ := app.Firestore(ctx)

	fmt.Println(client)
	return client, nil
}
