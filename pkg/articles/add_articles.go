package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/dto"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

func (h handler) AddArticle(c *gin.Context) {
	body := dto.AddArticleDTO{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := utils.GetTokenData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:    body.Title,
		Content:  body.Content,
		AuthorID: userID,
	}

	if err := h.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	articleDTO := dto.ArticleDTO{
		ID:       article.ID,
		Title:    article.Title,
		Content:  article.Content,
		AuthorID: article.AuthorID,
	}

	c.JSON(http.StatusCreated, &articleDTO)
}
