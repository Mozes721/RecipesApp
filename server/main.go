package main

import (
	"fmt"
	"github.com/RecepieApp/server/db"
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
type Map map[string]interface{}

func main() {
	fmt.Println(db.ReadCollection())
	recepie := Map{"Made": true, "Raiting": 7, "Title": "Hinkali"}

	db.AddCollectiosRecepie(recepie)
	fmt.Println(db.ReadCollection())
}
