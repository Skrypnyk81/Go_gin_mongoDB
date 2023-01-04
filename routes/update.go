package routes

import (
	"context"
	"fmt"
	"ginTest/Collection"
	"ginTest/databases"
	"ginTest/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = databases.ConnectDB()
	var postCollection = Collection.GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var post model.Posts

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"title": post.Title, "article": post.Article}
	//edited := bson.M{"article": post.Article}
	fmt.Println(post)
	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "data update successfully", "Data": res})
}
