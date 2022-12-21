package db

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Recepie struct {
	Made   bool
	Rating int
	Title  string
	Url    string
}

type HTTPError struct {
	Err      error
	Code     int
	FlashMsg string
	URL      string
}

func ReadCollection(ctx context.Context, client *firestore.Client) interface{} {
	iter := client.Collection("my-recepies").Documents(ctx)
	defer client.Close()
	var data []Recepie
	var recepie Recepie
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err := doc.DataTo(&recepie); err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data = append(data, recepie)
	}
	return data
}

func AddRecepie(client *firestore.Client, r Recepie) HTTPError {
	exists := checkIfExists(client, r.Title)
	if exists {
		return HTTPError{FlashMsg: "Title already exists"}
	} else {

		addCollectiosRecepie(context.Background(), client, r)
		return HTTPError{FlashMsg: ""}
	}
}

func UpdateRecepie(ctx context.Context, client *firestore.Client, recepie Recepie) error {
	uid := getRecepieID(client, recepie.Title)
	url := getUrlOfRecepie(client, recepie.Title)
	_, err := client.Collection("my-recepies").Doc(uid).Set(ctx, map[string]interface{}{
		"Made":   true,
		"Rating": recepie.Rating,
		"Title":  recepie.Title,
		"Url":    url,
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	return nil
}

func DeleteRecepie(ctx context.Context, client *firestore.Client, title string) error {
	uid := getRecepieID(client, title)
	_, err := client.Collection("my-recepies").Doc(uid).Delete(ctx)
	if err != nil {
		return err
		log.Printf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted recepie %v", err)
	return nil
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

func checkIfExists(client *firestore.Client, title string) bool {
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

func getUrlOfRecepie(client *firestore.Client, title string) string {
	ctx := context.Background()
	var r Recepie
	var url string
	iter := client.Collection("my-recepies").Where("Title", "==", title).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		if err := doc.DataTo(&r); err != nil {
			log.Fatalf("Failed to find recepie url: %v", err)
		}
		fmt.Println(url)
		url = r.Url
	}
	return url
}

func getRecepieID(client *firestore.Client, title string) string {
	ctx := context.Background()
	var uid string
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
			uid = doc.Ref.ID
		} else {
			uid = ""
		}
	}
	return uid
}
