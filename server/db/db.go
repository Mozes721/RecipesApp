package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"log"

	"google.golang.org/api/option"
)

func FirebaseDB() *firestore.Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile("/home/mozes/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	return client
}

func ReadCollection() (readCollection map[string]interface{}) {
	ctx := context.Background()
	client := FirebaseDB()
	defer client.Close()
	iter := client.Collection("my-recepies").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		readCollection = doc.Data()
	}
	return readCollection
}

func CheckCollection(recepie, collection map[string]string) {
	if checkForValue(recepie, ReadCollection()) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}
}

func checkForValue(recepie string, collection map[string]interface{}) bool {
	_, ok := collection["Title"][recepie]
	if ok {
		return true
	} else {
		return false
	}
}

func AddCollectiosRecepie(recepie map[string]interface{}) {
	ctx := context.Background()
	client := FirebaseDB()
	defer client.Close()
	_, _, err := client.Collection("my-recepies").Add(ctx, map[string]interface{}{
		"Made":   recepie["Made"],
		"Rating": recepie["Raiting"],
		"Title":  recepie["Title"],
		"Url":    recepie["Url"],
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
	return
}
