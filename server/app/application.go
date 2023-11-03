package app

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	. "github.com/RecepieApp/server/config"
)

type Application struct {
	FireClient *firestore.Client
	FireAuth *auth.Client
}

func (a *Application) LoadConfigurations() error {
	ctx := context.Background()

	fireClient, err := GetFirestoreClient(ctx)
	if err != nil {
        return err
    }
	a.FireClient = fireClient

	fireAuth, err := GetAuthClient(ctx)
	if err != nil {
        return err
    }
	a.FireAuth = fireAuth

	return nil
}
