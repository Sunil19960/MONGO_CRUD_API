package handlers

import (
	"crud/models"
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBook(c *gin.Context) {
	var book models.Book
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := utils.MongoClient.Database("books_db").Collection("book")
	filter := bson.M{"id": book.ID}
	update := bson.M{
		"$set": bson.M{
			"title":  book.Title,
			"author": book.Author,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "Failed",
			"Message": "Failed to add book",
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "Success",
		"Message": "Book updated Successfully",
		"Details": book,
	})
}
