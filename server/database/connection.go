package database

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func db() {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(app)
}
