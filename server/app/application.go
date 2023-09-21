package app

import (
	"context"

	"cloud.google.com/go/firestore"
	. "github.com/RecepieApp/server/config"
)

func LoadConfigurations() (*firestore.Client, error) {
	ctx := context.Background()
	client := FirebaseDB(ctx)
	if client == nil {
		return nil, nil
	}

	return client, nil
}
