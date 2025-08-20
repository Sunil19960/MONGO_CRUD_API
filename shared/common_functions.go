package shared

import (
	"context"
	"crud/utils"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertID() string {
	var insertId int64
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	collection := utils.MongoClient.Database("books_db").Collection("last_inserted_id")
	var result bson.M
	err := collection.FindOne(ctx, bson.M{}).Decode(&result)
	if err != nil {
		insertId = 1
	} else {
		last_inserted_id, err := strconv.ParseInt(result["last_inserted_id"].(string), 10, 64)

		if err != nil {

		}
		insertId = last_inserted_id + 1

	}
	return strconv.FormatInt(insertId, 10)
}
