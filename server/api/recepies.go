package api

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func showRecepies(ctx *gin.Context, client *firestore.Client) {
	userID := ctx.Query("userID")

	data, err := models.ReadUserCollection(ctx, client, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": "Unable to retrieve data",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)

}

func addRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Unable to parse data",
		})
		return
	}

	msg, status := data.AddRecepie(client)

	ctx.JSON(status, gin.H{
		"message": msg,
	})

}

func updateRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.RecepieUpdate{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Unable to parse data",
		})
		return
	}

	msg, status := data.UpdateRecepie(ctx, client)
	fmt.Println()
	ctx.JSON(status, gin.H{
		"message": msg,
	})
}

func deleteRecepie(ctx *gin.Context, client *firestore.Client) {
	recordID := ctx.Param("id")

	err, status := models.DeleteUserRecepie(ctx, client, recordID)
	if err != nil {
		ctx.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(status, gin.H{
		"message": "Recipe deleted successfully",
	})

}
