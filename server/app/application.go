package app

import (
	"context"
	"github.com/gobuffalo/envy"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	. "github.com/RecepieApp/server/config"
)

type Application struct {
	FireClient *firestore.Client
	FireAuth   *auth.Client
	ListenPort string
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

	a.ListenPort = envy.Get("PORT", "8080")

	return nil
}