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

func ReadCollection(ctx context.Context, client *firestore.Client) map[string]interface{} {
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
		return doc.Data()
	}
	return nil
}
func AddRecepie(client *firestore.Client, r Recepie) HTTPError {
	exists := checkIfExists(client, r.Title)
	if exists {
		fmt.Println("Already exists")
		return HTTPError{FlashMsg: "Title already exists"}
	} else {

		addCollectiosRecepie(context.Background(), client, r)
		return HTTPError{FlashMsg: ""}
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

func UpdateRecepie(ctx context.Context, client *firestore.Client, recepie Recepie) error {
	fmt.Println("Can add new recepie")
	uid := getRecepieID(client, recepie.Title)
	_, err := client.Collection("my-recepies").Doc(uid).Set(ctx, map[string]interface{}{
		"Made":   recepie.Made,
		"Rating": recepie.Rating,
		"Title":  recepie.Title,
		"Url":    recepie.Url,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	// [END firestore_data_set_field]
	return nil
}

func DeleteRecepie(ctx context.Context, client *firestore.Client, recepie Recepie) error {
	uid := getRecepieID(client, recepie.Title)
	_, err := client.Collection("my-recepies").Doc(uid).Delete(ctx)
	if err != nil {
		return err
		log.Printf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted recepie %v", err)
	return nil
}
