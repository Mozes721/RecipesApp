package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"time"
)

type UserCache struct {
	Authenticated bool   `redis:"Authenticated"`
	AuthToken     string `redis:"AuthToken"`
	UserID        string `redis:"UserID"`
	Email         string `redis:"Email"`
}

type Recepie struct {
	UserID   string
	Made     bool
	Rating   int
	Title    string
	Url      string
	ImageUrl string
}

func ReadUserCollection(ctx *gin.Context, client *firestore.Client, userID string) (*firestore.DocumentSnapshot, error) {
	projectID := "my-recepies"
	iter := client.Collection(projectID).Where("UserID", "==", userID).Documents(ctx)
	var data *firestore.DocumentSnapshot
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, err
		}

		data = doc
	}

	return data, nil
}

func (r *Recepie) AddRecepie(client *firestore.Client) (string, int) {
	var (
		message string
		status  int
	)
	exists := r.checkCollection(client)
	if exists {
		message = "Recepie is already added to your list of Yumms."
		status = 409
	} else {
		message = r.addCollectionRecepie(client)
		status = 200
	}

	return message, status
}

func (r *Recepie) addCollectionRecepie(client *firestore.Client) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := "Recepie successfully added to Yumms."

	_, _, err := client.Collection("my-recepies").Add(ctx, map[string]interface{}{
		"UserID":   r.UserID,
		"Made":     r.Made,
		"Rating":   r.Rating,
		"Title":    r.Title,
		"Url":      r.Url,
		"ImageUrl": r.ImageUrl,
	})
	if err != nil {
		message = "Issue with adding to your collection"
	}

	return message
}

func (r *Recepie) UpdateRecepie(ctx *gin.Context, client *firestore.Client) error {
	defer client.Close()
	exists := r.checkCollection(client)
	if !exists {
		return fmt.Errorf("failed updating document: %v", exists)
	}

	_, err := client.Collection("my-recepies").Doc(r.UserID+"_"+r.Title).Update(ctx, []firestore.Update{
		{
			Path:  "Made",
			Value: r.Made,
		},
		{
			Path:  "Rating",
			Value: r.Rating,
		},
	})
	if err != nil {
		ctx.JSON(400, gin.H{
			"Message": fmt.Errorf("Failed updating document: %v", err),
		})
	}

	return nil
}

func (r *Recepie) DeleteUserRecepie(ctx *gin.Context, client *firestore.Client) *firestore.DocumentSnapshot {
	projectID := "my-recepies"
	docRef := client.Collection(projectID).Doc(r.UserID + "_" + r.Title)
	snapshot, err := docRef.Get(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"Message": fmt.Errorf("Failed updating document: %v", err),
		})
	}
	if !snapshot.Exists() {
		ctx.JSON(404, gin.H{
			"Message": "Recipe not found",
		})

	}
	_, err = docRef.Delete(ctx)
	if err != nil {
		ctx.JSON(300, gin.H{
			"Message": fmt.Errorf("Failed to delete recipe: %v", err),
		})
	}

	return nil
}
