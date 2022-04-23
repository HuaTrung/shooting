package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"shootingplane/entity/api"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateArtistInput : struct for create art post request
type CreateArtistInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// FindArtists : Controller for getting all artists
func GetFriends(c *gin.Context) {
	client := c.MustGet("mongo").(*mongo.Client)
	friendCollection := client.Database("shooting").Collection("friend")
	var friends []api.Friend
	cursor, err := friendCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(context.TODO(), &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)

	c.JSON(http.StatusOK, gin.H{"data": friends})
}

// CreateArtist : controller for creating new artists
//func GetFriend(c *gin.Context) {
//	models := c.MustGet("models").(*mongo.Client)
//
//	// Validate input
//	var input CreateArtistInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Create artist
//	artist := Artist{Name: input.Name, Email: input.Email}
//	models.Create(&artist)
//
//	c.JSON(http.StatusOK, gin.H{"data": artist})
//}
