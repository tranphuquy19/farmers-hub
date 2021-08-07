package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tranphuquy19/farmers-hub/model"
	e "github.com/tranphuquy19/farmers-hub/pkg/error"
)

func Login(c *gin.Context) {
	var reqLogin model.LoginData

	if err := c.ShouldBindJSON(&reqLogin); err != nil {
		c.JSON(e.ErrorMessage(&e.RequestError{
			ErrorCode:   http.StatusBadRequest,
			ErrorString: err.Error(),
		}))
		c.Error(err).SetMeta("AuthHandler.Login")
		return
	}

	c.Set("data", true)
	c.JSON(http.StatusAccepted, reqLogin)
}

func Register(c *gin.Context) {
	var reqRegister model.RegisterData

	if err := c.ShouldBindJSON(&reqRegister); err != nil {
		c.JSON(e.ErrorMessage(&e.RequestError{
			ErrorCode:   http.StatusBadRequest,
			ErrorString: err.Error(),
		}))
		c.Error(err).SetMeta("AuthHandler.Register")
		return
	}

	c.JSON(http.StatusCreated, reqRegister)
}
