package handlers

import (
	"crud/models"
	"crud/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginReq models.LoginModel
	var token string
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Invalid Login Request! Please provide userId and password correctly",
		})
		return
	}

	if loginReq.UserId == "1234" && loginReq.Password == "1234" {
		var err error
		if token, err = shared.CreateJWTToken(loginReq.UserId); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  "failed",
				"message": "Internal Server Error",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Invalid UserUd or Password",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"token":   token,
		"message": "Authorization token has been generated successfully",
	})
}
