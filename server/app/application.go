package app

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
)

type Application struct {
	context   context.Context
	firestore *firestore.Client
}

var errStartLogger = errors.New("logger could not be created")

func (app *Application) Setup() error {
	fmt.Println("none")
	return nil
}
