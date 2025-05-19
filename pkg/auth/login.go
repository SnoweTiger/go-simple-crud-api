package auth

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/dto"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

// Login method
func (h handler) Login(c *gin.Context) {
	login := dto.LoginDTO{}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.LoginCheck(login.Login, login.Password, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	response := dto.LoginRespDTO{Login: login.Login, Token: token}
	c.JSON(http.StatusOK, &response)
}
