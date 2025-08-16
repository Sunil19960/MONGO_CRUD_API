package routes

import (
	"crud/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/status", handlers.GetEmployee)
	r.GET("/getEmployee", handlers.GetEmployee)
	r.POST("/postEmployee", handlers.GetEmployee)
	r.PUT("/updateEmployee", handlers.GetEmployee)
	r.DELETE("/deleteEmployee", handlers.GetEmployee)
}
