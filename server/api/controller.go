package api

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	md "github.com/RecepieApp/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, client *firestore.Client, auth *auth.Client) {
	router.Use(func(c *gin.Context) {
		md.AuthJWT(auth)(c)
	})

	router.GET("/", func(c *gin.Context) {
		showRecepies(c, client)
	})

	router.POST("/recipes", func(c *gin.Context) {
		addRecepie(c, client)
	})

	router.PUT("/recipes/:id", func(c *gin.Context) {
		updateRecepie(c, client)
	})
	router.DELETE("/recipes/:id", func(c *gin.Context) {
		deleteRecepie(c, client)
	})

	router.Run()
}
