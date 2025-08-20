package handlers

import (
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RemoveBook(c *gin.Context) {
	book_id := c.Param("id")

	collection := utils.MongoClient.Database("books_db").Collection("book")
	filter := bson.M{"id": book_id}

	ctx := c.Request.Context()

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "failed",
			"Message": "Book ID not found",
		})
		return
	}
	c.JSON(http.StatusOK,
		map[string]interface{}{
			"Status":  "Success",
			"Message": "Book removed Successfully",
		})

}
