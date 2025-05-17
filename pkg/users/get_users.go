package users

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(c *gin.Context) {
	var users []models.User
	var usersResponse []UserResponse

	if result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for _, user := range users {
		usersResponse = append(usersResponse, UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Login: user.Login,
		})
	}
	c.JSON(http.StatusOK, usersResponse)
}
