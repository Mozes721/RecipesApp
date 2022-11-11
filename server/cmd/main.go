package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RecepieApp/server/db"
	"github.com/RecepieApp/server/helpers"
)

//func server(ctx context) {
//	server := gin.Default()
//
//	server.GET("/test", func(ctx *gin.Context) {
//		ctx.JSON(200, gin.H{
//			"Message": "OK!",
//		})
//	})
//}

func main() {
	r := db.Recepie{Made: true, Rating: 7, Title: "Hinkali", Url: "https://hinkali.com"}
	// recepie := Map{"Made": true, "Raiting": 7, "Title": "Hinkali", "Url": "https://hinkali.com"}
	// CheckIfTitleExists(recepie["Title"])
	title := helpers.GetFieldTitle(&r, "Title")
	ctx := context.Background()
	client := db.FirebaseDB()
	recepies := db.ReadCollection(ctx, client)
	json_recepies, err := helpers.ToJson(recepies)
	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}
	titles := helpers.GetFieldDBTitles(json_recepies, "Title")
	fmt.Println(db.CheckIfRecepieExists(title, json_recepies))
	// AddCollectiosRecepie(recepie)
	// fmt.Println(ReadCollection(r))
}
