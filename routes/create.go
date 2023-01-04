package routes

import (
	"context"
	"ginTest/Collection"
	"ginTest/databases"
	"ginTest/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {
	var DB = databases.ConnectDB()
	var postCollection = Collection.GetCollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Posts)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := model.Posts{
		ID:      primitive.NewObjectID(),
		Title:   post.Title,
		Article: post.Article,
	}
	result, err := postCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}
