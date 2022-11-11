package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func FirebaseDB() *firestore.Client {
	ctx := context.Background()
	data, err := ioutil.ReadFile("C:\\Users\\RichardTaujenis\\Desktop\\RecipesApp\\server\\db\\serviceAccountKey.json")
	if err != nil {
		panic(err)
	}
	opt := option.WithCredentialsJSON(data)
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

func ReadCollection(ctx context.Context, client *firestore.Client) interface{} {
	defer client.Close()
	iter := client.Collection("my-recepies").Documents(ctx)
	defer iter.Stop()
	var data []Recepie
	var recepie Recepie
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err := doc.DataTo(&recepie); err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		data = append(data, recepie)
	}
	return data
}
func CheckIfRecepieExists(recepie string, recepies interface{}) bool {

	fmt.Println(recepie)
	return true
}

// func CheckIfTitleExists(ctx context.Context, title interface{}) bool {
// 	// [START get_user_by_email]
// 	collection := ReadCollection()
// 	u, err := CheckIfTitleExists(ctx, email)
// 	if err != nil {
// 		log.Fatalf("error getting user by email %s: %v\n", email, err)
// 	}
// 	log.Printf("Successfully fetched user data: %v\n", u)
// 	// [END get_user_by_email]
// 	return u
// }
// func CheckIfTitleExists(title interface{}) {
// 	data := ReadCollection()
// 	fmt.Println(data)

// }

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
