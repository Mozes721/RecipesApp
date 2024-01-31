package api

import (
	"cloud.google.com/go/firestore"
	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
)

func showRecepies(ctx *gin.Context, client *firestore.Client) {
	user := models.User{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &user)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	user.ReadUserCollection(ctx, client)
}

func addRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	msg, status := data.AddRecepie(ctx, client)

	ctx.JSON(status, gin.H{
			"message": msg,
		})
	
}

func updateRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	data.UpdateRecepie(ctx, client)

}

func deleteRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Message": "Unable to parse data",
		})
	}

	data.DeleteUserRecepie(ctx, client)

}
