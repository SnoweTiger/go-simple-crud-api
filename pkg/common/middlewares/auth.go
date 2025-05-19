package middlewares

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.TokenValid(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
