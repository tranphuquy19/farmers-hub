package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tranphuquy19/farmers-hub/dto"
	"github.com/tranphuquy19/farmers-hub/model"
	e "github.com/tranphuquy19/farmers-hub/pkg/error"
)

var jwtKey = []byte("abcdefghijklmnopq")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var reqLogin dto.LoginDTO

	if err := c.ShouldBindJSON(&reqLogin); err != nil {
		c.JSON(e.ErrorMessage(&e.RequestError{
			ErrorCode:   http.StatusBadRequest,
			ErrorString: err.Error(),
		}))
		c.Error(err).SetMeta("AuthHandler.Login")
		return
	}

	c.JSON(http.StatusAccepted, reqLogin)
}

func Register(c *gin.Context) {
	var reqRegister dto.RegisterDTO

	if err := c.ShouldBindJSON(&reqRegister); err != nil {
		c.JSON(e.ErrorMessage(&e.RequestError{
			ErrorCode:   http.StatusBadRequest,
			ErrorString: err.Error(),
		}))
		c.Error(err).SetMeta("AuthHandler.Register")
		return
	}

	user := model.User{
		ID:       uuid.New().String(),
		Username: reqRegister.Username,
		Password: reqRegister.Password,
	}

	c.JSON(http.StatusCreated, user)
}
