package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"github.com/RecepieApp/server/utils"
	"google.golang.org/api/iterator"
	"io"
)

func (r *Recepie) checkCollection(client *firestore.Client) bool {
	ctx := context.Background()
	iter := client.Collection(utils.ProjectID).Where("UserID", "==", r.UserID).Where("Title", "==", r.Title).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if doc != nil && doc.Exists() {
			return true
		}

	}

	return false
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
