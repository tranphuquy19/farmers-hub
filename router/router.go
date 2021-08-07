package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/handler"
)

func Routes(c *gin.Engine) {
	c.GET("/", handler.Index)

	c.GET("ping", handler.Ping)

	api := c.Group("api")
	{
		auth := api.Group("auth")
		{
			auth.POST("login", handler.Login)
			auth.POST("register", handler.Register)
		}
	}
}

func NoRoute(c *gin.Engine) {
	c.NoRoute(handler.ErrorHandler)
}
