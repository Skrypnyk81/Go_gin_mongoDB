package routes

import (
	"context"
	"ginTest/Collection"
	"ginTest/databases"
	"ginTest/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func ReadOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = databases.ConnectDB()
	var postCollection = Collection.GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var result model.Posts

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)
	err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
