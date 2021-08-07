package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/config"
	r "github.com/tranphuquy19/farmers-hub/router"
)

func main() {
	router := gin.Default()

	r.NoRoute(router)
	r.Routes(router)

	router.Run(config.APP_PORT) // listen and serve on 0.0.0.0:8080
}
