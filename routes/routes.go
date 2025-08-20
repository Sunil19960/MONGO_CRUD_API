package routes

import (
	"context"
	"crud/handlers"
	"crud/utils"

	"github.com/gin-gonic/gin"
)

func setContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.Background())
		c.Next()
	}
}

func InitializeRoutes(r *gin.Engine) {
	r.GET(utils.Settings.GetStatus, setContext(), handlers.GetStatus)
	r.GET(utils.Settings.GetBookRoute, setContext(), handlers.GetBook)
	r.POST(utils.Settings.AddBookRoute, setContext(), handlers.AddBook)
	r.POST(utils.Settings.UpdateBookRoute, setContext(), handlers.UpdateBook)
	r.DELETE(utils.Settings.RemoveBookRoute, setContext(), handlers.RemoveBook)
}
