package users

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ChangePassRequest struct {
	ID          int    `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (h handler) ChangePass(c *gin.Context) {
	body := ChangePassRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if result := h.DB.First(&user, body.ID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	err := VerifyPassword(body.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user.Password = string(hashedNewPassword)

	if err := h.DB.Save(&user).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusAccepted)
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
