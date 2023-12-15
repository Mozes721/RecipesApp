package models

import (
	"cloud.google.com/go/firestore"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

type User struct {
	UserID string
}

type Recepie struct {
	User
	Made   bool
	Rating int
	Title  string
	Url    string
}

func (u *User) ReadUserCollection(c *gin.Context, client *firestore.Client) *firestore.DocumentSnapshot {
	projectID := "my-recepies"
	iter := client.Collection(projectID).Where("UserID", "==", u.UserID).Documents(c)
	var data *firestore.DocumentSnapshot
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			c.JSON(404, gin.H{
				"Message": "Failed to iterate",
			})
		}

		data = doc
	}

	return data
}

func (r *Recepie) AddRecepie(c *gin.Context, client *firestore.Client) error {
	ok := r.checkCollection(client)
	if ok {
		c.JSON(200, gin.H{
			"Message": "Recepie is already added to your list of Yumms.",
		})
	} else {
		err := r.addCollectionRecepie(c, client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Recepie) addCollectionRecepie(c *gin.Context, client *firestore.Client) error {
	defer client.Close()
	_, _, err := client.Collection("my-recepies").Add(c, map[string]interface{}{
		"UserID": r.User.UserID,
		"Made":   r.Made,
		"Rating": r.Rating,
		"Title":  r.Title,
		"Url":    r.Url,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"Message": "Recepie is already added to your list of Yumms",
		})
	}

	return nil
}

func (r *Recepie) UpdateRecepie(c *gin.Context, client *firestore.Client) error {
	exists := r.checkCollection(client)
	if !exists {
		return fmt.Errorf("failed updating document: %v", exists)
	}

	_, err := client.Collection("my-recepies").Doc(r.User.UserID+"_"+r.Title).Update(c, []firestore.Update{
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
	projectID := "my-recepies"
	docRef := client.Collection(projectID).Doc(r.User.UserID + "_" + r.Title)
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
