package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/handler"
)

func Routes(c *gin.Engine) {
	c.GET("/login", handler.Login)
	c.GET("/ping", handler.Ping)
}

func NoRoute(c *gin.Engine) {
	c.NoRoute(handler.ErrorHandler)
}
