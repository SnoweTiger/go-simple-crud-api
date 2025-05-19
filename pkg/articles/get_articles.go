package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/dto"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetArticles(c *gin.Context) {
	var articles []models.Article

	if err := h.DB.Model(&models.Article{}).Preload("Author").Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var articlesDTO []dto.ArticleDTO
	for _, a := range articles {
		articlesDTO = append(articlesDTO, dto.ArticleDTO{
			ID:      a.ID,
			Title:   a.Title,
			Content: a.Content,
			Author: dto.AuthorDTO{
				ID:   a.Author.ID,
				Name: a.Author.Name,
			},
		})
	}

	c.JSON(http.StatusOK, articlesDTO)
}
