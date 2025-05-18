package articles

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddArticleRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId uint   `json:"authorId"`
}

func (h handler) AddArticle(c *gin.Context) {
	body := AddArticleRequest{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// userID, err := utils.GetTokenData(c)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// fmt.Printf("userId = %d", userID)

	var article models.Article
	article.Title = body.Title
	article.Content = body.Content
	article.AuthorId = body.AuthorId

	if err := h.DB.Create(&article).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, &article)
}
