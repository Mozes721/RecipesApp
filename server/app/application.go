package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	. "github.com/RecepieApp/server/config"
)

type Application struct {
	context         context.Context
	firestore       *firestore.Client
	Client          *firestore.Client
	ControlDBConfig string
}

func (a *Application) LoadConfigurations() (*firestore.Client, error) {
	fmt.Println("asdsd")
	client, err := FirebaseDB(a.context)
	if err != nil {
		panic("could not connect to firebase client")
	}
	return client, nil
}
