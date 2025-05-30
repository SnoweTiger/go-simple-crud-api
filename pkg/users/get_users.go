package users

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/dto"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(c *gin.Context) {
	var users []models.User
	var usersResponse []dto.UserDTO

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range users {
		usersResponse = append(usersResponse, dto.UserDTO{
			ID:    user.ID,
			Name:  user.Name,
			Login: user.Login,
		})
	}
	c.JSON(http.StatusOK, usersResponse)
}
