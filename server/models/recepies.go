package models

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)



func ReadCollection(ctx context.Context, client *firestore.Client) *firestore.DocumentSnapshot {
	projectID := "my-recepies"
	iter := client.Collection(projectID).Documents(ctx)
	var data *firestore.DocumentSnapshot
	auth := client.AuthService()
	auth := client.FireAuth()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		data = doc
	}

	return data
}

func AddRecepie(client *firestore.Client, r Recepie) error {
	ok := checkCollection(client, r.Title)
	if ok {
		fmt.Println("Apready exists")
		return fmt.Errorf({FlashMsg: "Title already exists"})
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

func addCollectiosRecepie(ctx context.Context, client *firestore.Client, recepie helpers.Recepie) {
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

func UpdateRecepie(ctx context.Context, client *firestore.Client, recepie Recepie) error {
	// First, check if the document with the specified title exists
	exists := checkCollection(client, recepie.Title)
	if !exists {
		return fmt.Errorf("Recipe with Title '%s' not found", recepie.Title)
	}

	_, err := client.Collection("my-recepies").Doc(recepie.Title).Update(ctx, []firestore.Update{
		{
			Path:  "Made",
			Value: recepie.Made,
		},
		{
			Path:  "Rating",
			Value: recepie.Rating,
		},
		// Add other fields you want to update
	})
	if err != nil {
		return fmt.Errorf("Failed updating document: %v", err)
	}

	return nil
}
