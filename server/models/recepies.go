package models

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/RecepieApp/server/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
	"time"
)

type Recepies struct {
	Recepies []Recepie `json:"recepies"`
}

type Recepie struct {
	ID       string `json:"id"`
	UserID   string `json:"userID"`
	Made     bool   `json:"made"`
	Rating   int    `json:"rating"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	ImageUrl string `json:"imageUrl"`
}

type RecepieUpdate struct {
	ID     string `json:"id"`
	Rating int    `json:"rating"`
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
		recepie.ID = doc.Ref.ID

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
		status = http.StatusConflict
	} else {
		message = r.addCollectionRecepie(client)
		status = http.StatusOK
	}

	return message, status
}

func (r *Recepie) addCollectionRecepie(client *firestore.Client) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := "Recepie successfully added to Yumms."

	_, _, err := client.Collection(utils.ProjectID).Add(ctx, map[string]interface{}{
		"UserID":   r.UserID,
		"Made":     false,
		"Rating":   0,
		"Title":    r.Title,
		"Url":      r.Url,
		"ImageUrl": r.ImageUrl,
	})
	if err != nil {
		message = "Issue with adding to your collection"
	}

	return message
}

func (ru *RecepieUpdate) UpdateRecepie(c *gin.Context, client *firestore.Client) (string, int) {
	docRef := client.Collection(utils.ProjectID).Doc(ru.ID)
	_, err := docRef.Update(c.Request.Context(), []firestore.Update{
		{Path: "Made", Value: true},
		{Path: "Rating", Value: ru.Rating},
	})
	if err != nil {
		return fmt.Sprintf("failed to update recipe: %v", err), http.StatusInternalServerError
	}

	return "Recipe updated successfully", http.StatusOK
}

func DeleteUserRecepie(c *gin.Context, client *firestore.Client, recordID string) (error, int) {
	docRef := client.Collection(utils.ProjectID).Doc(recordID)
	snapshot, err := docRef.Get(c)
	if err != nil {
		return fmt.Errorf("failed updating document: %v", err), http.StatusInternalServerError

	}
	if !snapshot.Exists() {
		return fmt.Errorf("recipe not found"), http.StatusNotFound

	}
	_, err = docRef.Delete(c)
	if err != nil {
		return fmt.Errorf("failed to delete recipe: %v", err), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
