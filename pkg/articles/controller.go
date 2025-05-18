package articles

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/articles")
	routes.GET("/", h.GetArticles)
	routes.POST("/searchByAuthor", h.GetArticlesByAuthor)
	routes.POST("/", h.AddArticle)
	routes.PUT("/", h.UpdateArticles)
	routes.DELETE("/:id", h.DeleteArticle)

}
