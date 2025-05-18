package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type ArticlesRequest struct {
	AuthorId uint `json:"authorId"`
}

func (h handler) GetArticles(c *gin.Context) {
	var articles []models.Article

	if result := h.DB.Find(&articles); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, articles)
}

func (h handler) GetArticlesByAuthor(c *gin.Context) {
	body := ArticlesRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var articles []models.Article

	if err := h.DB.Find(&articles, "author_id = ?", body.AuthorId).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, articles)
}
