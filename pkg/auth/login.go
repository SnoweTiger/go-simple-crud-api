package auth

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Login    string `json:"login" `
	Password string `json:"password"`
}

type LoginResponse struct {
	Login string `json:"login"`
	Token string `json:"token"`
}

// Login method
func (h handler) Login(c *gin.Context) {
	login := LoginRequest{}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := utils.LoginCheck(login.Login, login.Password, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	response := LoginResponse{Login: login.Login, Token: token}
	c.JSON(http.StatusOK, &response)
}
