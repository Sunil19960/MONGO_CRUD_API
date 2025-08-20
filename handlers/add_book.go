package handlers

import (
	"crud/models"
	"crud/shared"
	"crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddBook(c *gin.Context) {
	var book models.Book
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := utils.MongoClient.Database("books_db").Collection("book")
	book.ID = shared.InsertID()

	_, err := collection.InsertOne(ctx, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "Failed",
			"Message": "Failed to add book",
		})
		return
	}
	idCollection := utils.MongoClient.Database("books_db").Collection("last_inserted_id")
	update := bson.M{
		"$set": bson.M{
			"last_inserted_id": book.ID,
		},
	}
	opts := options.Update().SetUpsert(true)
	idCollection.UpdateOne(ctx, bson.M{}, update, opts)
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "Success",
		"Message": "Book inserted Successfully",
		"Details": book,
	})
}
