package handlers

import (
	"crud/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	status, _ := database.PingMongoDB()
	if status == "Success" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  "Success",
			"Message": "MongoDB Connection Successful",
		})
	}
}
