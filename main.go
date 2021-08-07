package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/config"
	"github.com/tranphuquy19/farmers-hub/handler"
)

func main()  {
	router := gin.Default()

	router.Use(handler.ErrorHandler)
	api := router.Group("/api")
	{
		api.GET("/ping", handler.Ping)
	}

	router.Run(config.APP_PORT) // listen and serve on 0.0.0.0:8080
}
