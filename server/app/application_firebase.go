package app

import (
	"cloud.google.com/go/firestore"
	"context"
	. "github.com/RecepieApp/server/config"
)

func LoadFirebaseConnection() (*firestore.Client, error) {
	ctx := context.Background()
	client := FirebaseDB(ctx)
	if client == nil {
		return nil, nil
	}

	return client, nil
}
