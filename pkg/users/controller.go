package users

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/users")
	routes.GET("/", h.GetUsers)
	routes.GET("/:id", h.GetUser)
	routes.POST("/", h.AddUser)
	routes.PUT("/", h.UpdateUser)
	routes.PATCH("/changepass", h.ChangePass)
	routes.DELETE("/:id", h.DeleteUser)
}
