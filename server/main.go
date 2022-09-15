package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"

	"github.com/RecepieApp/server/database"
)

// func connect(app) {
// 	client, err := app.Firestore(ctx)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func main() {
	ctx := context.Background()
	// server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"Message": "OK!",
	// 	})
	// })
	db := database.FirebaseDB()
	// readDb := database.ReadData()

	client, err := db.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(readDb)
	fmt.Println(client)

}
