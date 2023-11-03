package api

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, client *firestore.Client) {
	
	router.POST("/login", func(c *gin.Context) {

	})

	router.POST("/register", func(c *gin.Context) {
		
	})

	router.POST("/googleLogin", func(c *gin.Context) {
		
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
