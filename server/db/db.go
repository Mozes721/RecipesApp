package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func FirebaseDB(ctx context.Context) *firestore.Client {
	opt := option.WithCredentialsFile("/home/mozes/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	return client
}

func ReadCollection(ctx context.Context) {
	ct := context.Background()
	projectID := "my-recepies"
	client := FirebaseDB(ct)
	iter := client.Collection(projectID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
}
func AddRecepie(client *firestore.Client, r *Recepie, title string) HTTPError {
	ok := checkCollection(client, title)
	if ok {
		fmt.Println("Apready exists")
		return HTTPError{FlashMsg: "Title already exists"}
	} else {
		fmt.Println("Can add new recepie")
		addCollectiosRecepie(r)
		return HTTPError{FlashMsg: ""}
	}
}

func checkCollection(client *firestore.Client, title string) bool {
	ctx := context.Background()
	var exists bool
	iter := client.Collection("my-recepies").Where("Title", "==", title).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		if doc.Data() != nil {
			exists = true
		} else {
			exists = false
		}
	}
	return exists

}

func addCollectiosRecepie(recepie *Recepie) {
	ctx := context.Background()
	client := FirebaseDB(ctx)
	defer client.Close()
	_, _, err := client.Collection("my-recepies").Add(ctx, map[string]interface{}{
		"Made":   recepie.Made,
		"Rating": recepie.Rating,
		"Title":  recepie.Title,
		"Url":    recepie.Url,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}
