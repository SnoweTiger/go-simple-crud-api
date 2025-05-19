package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if err := h.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.DB.Delete(&article)

	c.Status(http.StatusOK)
}
