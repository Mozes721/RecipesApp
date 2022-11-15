package main

import (
	"context"
	"fmt"

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
	r := db.Recepie{Made: true, Rating: 7, Title: "Spagetti", Url: "https://hinkali.com"}

	// recepie := Map{"Made": true, "Raiting": 7, "Title": "Hinkali", "Url": "https://hinkali.com"}
	// CheckIfTitleExists(recepie["Title"])
	ctx := context.Background()
	client := db.FirebaseDB(ctx)
	title := helpers.GetFieldTitle(&r, "Title")
	setToMap := helpers.SetToMap(&r)
	db.AddRecepie(ctx, client, setToMap, title)

	fmt.Println(title)
	// AddCollectiosRecepie(recepie)
	db.ReadCollection(ctx)
}
