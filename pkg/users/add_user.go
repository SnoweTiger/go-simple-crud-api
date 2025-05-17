package users

import (
	"html"
	"net/http"
	"strings"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type AddUserRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	user.Name = body.Name

	user.Login = html.EscapeString(strings.TrimSpace(body.Login))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user.Password = string(hashedPassword)

	if err := h.DB.Create(&user).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userResponse := UserResponse{
		ID:    user.ID,
		Login: user.Login,
		Name:  user.Name,
	}

	c.JSON(http.StatusCreated, &userResponse)
}
