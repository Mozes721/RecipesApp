package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func ReadCollection(ctx context.Context, client *firestore.Client) {
	projectID := "my-recepies"
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
func AddRecepie(client *firestore.Client, r Recepie) HTTPError {
	ok := checkCollection(client, r.Title)
	if ok {
		fmt.Println("Apready exists")
		return HTTPError{FlashMsg: "Title already exists"}
	} else {
		fmt.Println("Can add new recepie")
		addCollectiosRecepie(context.Background(), client, r)
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

func addCollectiosRecepie(ctx context.Context, client *firestore.Client, recepie Recepie) {
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
