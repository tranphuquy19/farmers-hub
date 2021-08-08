package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tranphuquy19/farmers-hub/dto"
	"github.com/tranphuquy19/farmers-hub/model"
	e "github.com/tranphuquy19/farmers-hub/pkg/error"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user := model.User{
		ID:       uuid.New().String(),
		Username: reqRegister.Username,
		Password: string(hashedPassword),
	}

	c.JSON(http.StatusCreated, user)
}
