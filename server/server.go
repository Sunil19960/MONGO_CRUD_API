package server

import (
	"crud/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	routes.InitializeRoutes(router)
	router.Run(":3000")
}
