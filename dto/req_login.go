package dto

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
