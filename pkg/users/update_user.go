package users

import (
	"html"
	"net/http"
	"strings"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/dto"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

func (h handler) UpdateUser(c *gin.Context) {
	body := models.User{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	user.ID = body.ID
	user.Name = body.Name

	user.Login = html.EscapeString(strings.TrimSpace(body.Login))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Password = string(hashedPassword)

	if err := h.DB.Save(&user).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	userResponse := dto.UserDTO{
		ID:    user.ID,
		Login: user.Login,
		Name:  user.Name,
	}

	c.JSON(http.StatusOK, &userResponse)
}
