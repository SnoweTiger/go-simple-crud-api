package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if result := h.DB.First(&article, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&article)

	c.Status(http.StatusOK)
}
