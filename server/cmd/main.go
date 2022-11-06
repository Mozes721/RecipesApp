package main

import (
	"fmt"

	. "github.com/RecepieApp/server/db"
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
	// recepie := Map{"Made": true, "Raiting": 7, "Title": "Hinkali"}
	// CheckIfTitleExists(recepie["Title"])
	fmt.Println(ReadCollection())
	// AddCollectiosRecepie(recepie)
	// fmt.Println(ReadCollection(r))
}
