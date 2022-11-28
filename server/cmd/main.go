package main

import (
	"flag"
	"github.com/RecepieApp/server/app"
	. "github.com/RecepieApp/server/runtime"
	//
	//"github.com/RecepieApp/server/db"
	//"github.com/RecepieApp/server/helpers"
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

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()
	// Initialize the runtime
	a := app.Application{}
	client, err := a.LoadConfigurations()
	if err != nil {
		panic(err)
	}
	err = Start(client)
	if err != nil {
		panic(err)
	}
}
