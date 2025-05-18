package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateArticles(c *gin.Context) {
	article := models.Article{}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.DB.Save(&article).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &article)
}
