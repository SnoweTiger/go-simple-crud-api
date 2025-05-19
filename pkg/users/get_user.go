package users

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/dto"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := h.DB.Model(&models.User{}).Preload("Articles").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var articlesDTO []dto.ArticleDTO
	for _, a := range user.Articles {
		articlesDTO = append(articlesDTO, dto.ArticleDTO{
			ID:    a.ID,
			Title: a.Title,
		})
	}

	data := dto.UserDataDTO{
		User: dto.UserDTO{
			ID:   user.ID,
			Name: user.Name,
		},
		Articles: articlesDTO,
	}

	c.JSON(http.StatusOK, &data)
}
