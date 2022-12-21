package api

import (
	"context"
	"fmt"
	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func allRecepies(c *gin.Context) {
	ctx := context.Background()
	client, err := app.LoadFirebaseConnection()
	if err != nil {
		err = fmt.Errorf("error occurred while loading")
	}

	defer client.Close()
	values := db.ReadCollection(ctx, client)
	c.JSON(200, values)
}

func newRecepie(c *gin.Context) {
	client, err := app.LoadFirebaseConnection()
	if err != nil {
		err = fmt.Errorf("error occurred while loading")
	}

	defer client.Close()
	var recepie db.Recepie
	if err := c.BindJSON(&recepie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, db.AddRecepie(client, recepie))
}

func deleteRecepie(c *gin.Context) {
	ctx := context.Background()
	client, err := app.LoadFirebaseConnection()
	if err != nil {
		err = fmt.Errorf("error occurred while ")
	}

	defer client.Close()
	var recepie db.Recepie
	if err := c.BindJSON(&recepie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, db.DeleteRecepie(ctx, client, recepie.Title))
}

func updateRecepie(c *gin.Context) {
	ctx := context.Background()
	client, err := app.LoadFirebaseConnection()
	if err != nil {
		err = fmt.Errorf("error occurred while ")
	}

	defer client.Close()
	var recepie db.Recepie
	if err := c.BindJSON(&recepie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, db.UpdateRecepie(ctx, client, recepie))
}
