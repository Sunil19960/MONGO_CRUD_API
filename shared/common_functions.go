package shared

import (
	"context"
	"crud/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

var secretKey = []byte("secret-key")

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

func CreateJWTToken(userID string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userID,
			"exp":      time.Now().Add(time.Minute * 15).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = tokenString[len("Bearer "):]
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Authorization Header missing",
			})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Invalid Authorization Header. Please Login again",
			})
			return
		}
		c.Next()
	}
}
