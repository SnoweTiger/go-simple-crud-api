package utils

import (
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(login string, password string, DB *gorm.DB) (string, error) {

	var err error

	u := models.User{}

	if err := DB.Model(models.User{}).Where("login = ?", login).Take(&u).Error; err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
