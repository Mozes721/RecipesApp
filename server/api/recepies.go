package api

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func showRecepies(ctx *gin.Context, client *firestore.Client) {
	ctx.JSON(200, gin.H{
		"Message": "recepies",
	})
}

func addRecepie(ctx *gin.Context, client *firestore.Client) {
	ctx.JSON(200, gin.H{
		"Message": "recepies add",
	})
}

func updateRecepie(ctx *gin.Context, client *firestore.Client) {
	ctx.JSON(200, gin.H{
		"Message": "recepies update",
	})

}

func deleteRecepie(ctx *gin.Context, client *firestore.Client) {
	ctx.JSON(200, gin.H{
		"Message": "recepies deleted",
	})

}
