package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/RecepieApp/server/database"
)

func main() {

	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "OK!",
		})
	})

	fmt.Println(database.Test("Rico"))

}
