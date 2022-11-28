package app

import (
	"cloud.google.com/go/firestore"
	"context"
	. "github.com/RecepieApp/server/config"
)

type Application struct {
	firestore       *firestore.Client
	Client          *firestore.Client
	ControlDBConfig string
}

func (a *Application) LoadConfigurations() (*firestore.Client, error) {
	ctx := context.Background()
	client := FirebaseDB(ctx)
	if client == nil {
		return nil, nil
	}

	return client, nil
}
