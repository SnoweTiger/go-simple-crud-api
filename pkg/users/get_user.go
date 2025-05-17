package users

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	userResponse := UserResponse{
		ID:    user.ID,
		Login: user.Login,
		Name:  user.Name,
	}

	c.JSON(http.StatusCreated, &userResponse)
}
