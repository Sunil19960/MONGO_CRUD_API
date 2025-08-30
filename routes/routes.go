package routes

import (
	"context"
	"crud/handlers"
	"crud/shared"
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
	r.POST(utils.Settings.LoginRoute, handlers.LoginHandler)

	r.GET(utils.Settings.GetStatus, shared.VerifyToken(), setContext(), handlers.GetStatus)
	r.GET(utils.Settings.GetBookRoute, shared.VerifyToken(), setContext(), handlers.GetBook)
	r.POST(utils.Settings.AddBookRoute, shared.VerifyToken(), setContext(), handlers.AddBook)
	r.POST(utils.Settings.UpdateBookRoute, shared.VerifyToken(), setContext(), handlers.UpdateBook)
	r.DELETE(utils.Settings.RemoveBookRoute, shared.VerifyToken(), setContext(), handlers.RemoveBook)

}
