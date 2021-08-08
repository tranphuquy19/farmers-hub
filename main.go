package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/config"
	"github.com/tranphuquy19/farmers-hub/pkg/cors"
	r "github.com/tranphuquy19/farmers-hub/router"
)

func main() {
	router := gin.Default()

	router.Static("/res", "public")
	router.LoadHTMLGlob("views/*")
	router.Use(cors.CORS())

	r.NoRoute(router)
	r.Routes(router)

	// gin.SetMode(gin.ReleaseMode) // Production mode
	router.Run(config.APP_PORT) // listen and serve on 0.0.0.0:8080
}
