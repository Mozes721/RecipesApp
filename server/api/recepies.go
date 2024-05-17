package api

import (
	"cloud.google.com/go/firestore"
	"fmt"
	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func showRecepies(ctx *gin.Context, client *firestore.Client) {
	user := models.User{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse data",
		})
		return
	}

	data, err := models.ReadUserCollection(ctx, client, user.UserID)
	if err != nil {
		ctx.JSON(404, gin.H{
			"Message": data,
		})
	}
	fmt.Println(data.Data())
	ctx.JSON(404, gin.H{
		"Message": "Failed to iterate",
	})

}

func addRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	msg, status := data.AddRecepie(client)

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
