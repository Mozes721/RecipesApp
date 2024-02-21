package api

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	md "github.com/RecepieApp/server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(router *gin.Engine, client *firestore.Client, auth *auth.Client) {
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Use(func(c *gin.Context) {
		md.AuthJWT(auth)(c)
	})

	router.GET("/:id", func(c *gin.Context) {
		userID := c.Param("id")

		showRecepies(c, client, userID)
	})

	router.POST("/recipe", func(c *gin.Context) {
		addRecepie(c, client)
	})

	router.PATCH("/recipe/:id", func(c *gin.Context) {
		updateRecepie(c, client)
	})
	router.DELETE("/recipes/:id", func(c *gin.Context) {
		deleteRecepie(c, client)
	})

	router.Run()
}
