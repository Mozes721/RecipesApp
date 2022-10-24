package main

import (
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
	//recepie := Map{"Made": false, "Raiting": 4, "Title": "Spagetti", "Url": "https://spagetti.com"}
	//db.AddCollectiosRecepie(recepie)
	test := "Spagetti"
	db.CheckCollection(test)
	//fmt.Println(db.CheckCollection(recepie))
}
