package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/config"
)

func main() {
	router := gin.Default()

	NoRoute(router)
	Routes(router)

	router.Run(config.APP_PORT) // listen and serve on 0.0.0.0:8080
}
