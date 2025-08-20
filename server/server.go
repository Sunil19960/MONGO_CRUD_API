package server

import (
	"context"
	"crud/config"
	"crud/database"
	"crud/routes"
	"crud/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	var port_no string
	utils.Settings = config.LoadSettings()
	database.ConnectMongoDB()
	defer utils.MongoClient.Disconnect(context.Background())
	router := gin.Default()
	routes.InitializeRoutes(router)
	port_no = fmt.Sprintf(":%s", utils.Settings.PortNo)
	router.Run(port_no)
}
