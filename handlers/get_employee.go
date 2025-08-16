package handlers

import (
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, models.Employee{
		ID:   1,
		Name: "Ramesh",
	})
}
