package handlers

import (
	"crud/models"
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBook(c *gin.Context) {
	book_id := c.Param("id")

	collection := utils.MongoClient.Database("books_db").Collection("book")
	filter := bson.M{"id": book_id}

	var book models.Book
	ctx := c.Request.Context()

	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "Failed",
			"Message": "Book not found",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "Success",
		"Details": book,
	})
}
