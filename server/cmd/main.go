package main

import (
	"context"
	"fmt"

	"github.com/RecepieApp/server/db"
	"github.com/RecepieApp/server/helpers"
)

//
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
	r := db.Recepie{Made: false, Rating: 0, Title: "Fish And Chips", Url: "https://www.thespruceeats.com/best-fish-and-chips-recipe-434856"}

	// recepie := Map{"Made": true, "Raiting": 7, "Title": "Hinkali", "Url": "https://hinkali.com"}
	// CheckIfTitleExists(recepie["Title"])
	ctx := context.Background()
	client := db.FirebaseDB(ctx)
	title := helpers.GetFieldTitle(&r, "Title")

	db.AddRecepie(client, &r, title)

	fmt.Println(title)
	// AddCollectiosRecepie(recepie)
	db.ReadCollection(ctx)
}
