package handler

import (
	"github.com/gin-gonic/gin"
	e "github.com/tranphuquy19/farmers-hub/pkg/error"
)

func ErrorHandler(c *gin.Context) {
	c.JSON(e.ErrorMessage(e.ErrNotFound))
}
