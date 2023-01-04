package routes

import (
	"context"
	"ginTest/Collection"
	"ginTest/databases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func DeletePost(c *gin.Context) {
	ctx, candel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = databases.ConnectDB()
	postId := c.Param("postId")

	var postCollection = Collection.GetCollection(DB, "Posts")
	defer candel()
	objId, _ := primitive.ObjectIDFromHex(postId)
	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})
}
