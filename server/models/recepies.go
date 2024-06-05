package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/RecepieApp/server/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"time"
)

type Recepies struct {
	Recepies []Recepie `json:"recepies"`
}

type Recepie struct {
	UserID   string `json:"userID"`
	Made     bool   `json:"made"`
	Rating   int    `json:"rating"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	ImageUrl string `json:"imageUrl"`
}

func ReadUserCollection(c *gin.Context, client *firestore.Client, userID string) (Recepies, error) {
	iter := client.Collection(utils.ProjectID).Where("UserID", "==", userID).Documents(c)
	var recepies Recepies

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return recepies, err
		}

		var recepie Recepie
		err = doc.DataTo(&recepie)
		if err != nil {
			fmt.Printf("Unable to unmarshal Firestore document data to struct due to %s\n", err)
			continue
		}
		recepies.Recepies = append(recepies.Recepies, recepie)
	}

	return recepies, nil

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

	_, _, err := client.Collection(utils.ProjectID).Add(ctx, map[string]interface{}{
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

func (r *Recepie) UpdateRecepie(c *gin.Context, client *firestore.Client) error {
	defer client.Close()
	exists := r.checkCollection(client)
	if !exists {
		return fmt.Errorf("failed updating document: %v", exists)
	}

	_, err := client.Collection(utils.ProjectID).Doc(r.UserID+"_"+r.Title).Update(c, []firestore.Update{
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
		c.JSON(400, gin.H{
			"Message": fmt.Errorf("Failed updating document: %v", err),
		})
	}

	return nil
}

func (r *Recepie) DeleteUserRecepie(c *gin.Context, client *firestore.Client) *firestore.DocumentSnapshot {
	docRef := client.Collection(utils.ProjectID).Doc(r.UserID + "_" + r.Title)
	snapshot, err := docRef.Get(c)
	if err != nil {
		c.JSON(500, gin.H{
			"Message": fmt.Errorf("Failed updating document: %v", err),
		})
	}
	if !snapshot.Exists() {
		c.JSON(404, gin.H{
			"Message": "Recipe not found",
		})

	}
	_, err = docRef.Delete(c)
	if err != nil {
		c.JSON(300, gin.H{
			"Message": fmt.Errorf("Failed to delete recipe: %v", err),
		})
	}

	return nil
}
