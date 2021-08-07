package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/model"
)

func Login(c *gin.Context) {
	var reqLogin model.LoginData

	if err := c.BindJSON(&reqLogin); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, reqLogin)
}

func Register(c *gin.Context) {
	var reqLogin model.LoginData

	if err := c.BindJSON(&reqLogin); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, reqLogin)
}
