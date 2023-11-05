package api

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, client *firestore.Client, auth *auth.Client) {
	//router.Use(middleware.AuthJWT(auth))

	router.POST("/login", func(c *gin.Context) {
		login(c, auth)
	})

	router.POST("/register", func(c *gin.Context) {
		register(c, auth)
	})

	router.POST("/googleLogin", func(c *gin.Context) {
		googleLogin(c, auth)
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
