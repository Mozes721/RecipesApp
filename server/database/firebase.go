package database

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func FirebaseDB() *firebase.App {
	ctx := context.Background()
	opt := option.WithCredentialsFile("/home/mozes/serviceAccountKey.json")
	config := &firebase.Config{ProjectID: "Recepies"}

	app, err := firebase.NewApp(ctx, config, opt)

	// client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	// defer client.Close()

	return app
}

func ReadData() *firebase.App {
	ctx := context.Background()
	app := FirebaseDB()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := client.Collection("users").Documents(ctx)
	for {
		doc, _ := iter.Next()
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	// return (*firebase.App)(iter)
	// defer client.Close()
}

