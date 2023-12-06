package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/api/iterator"
	"io"
	"log"
)

func (r *Recepie) checkCollection(client *firestore.Client) bool {
	ctx := context.Background()
	var exists bool
	iter := client.Collection("my-recepies").Where("UserID", "==", r.User.userID).Where("Title", "==", r.Title).Documents(ctx)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
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

func UnmarshallRequestBodyToAPIData(requestBody io.ReadCloser, respStruct interface{}) error {
	defer requestBody.Close()

	body, err := io.ReadAll(requestBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, respStruct)
	if err != nil {
		return err
	}

	return nil
}
