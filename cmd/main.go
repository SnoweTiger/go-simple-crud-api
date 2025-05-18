package main

import (
	"net/http"

	"github.com/SnoweTiger/go-simple-crud-api/pkg/articles"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/auth"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/db"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/common/middlewares"
	"github.com/SnoweTiger/go-simple-crud-api/pkg/users"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	auth.RegisterRoutes(api, h)

	articles.RegisterRoutes(api, h)

	api.Use(middlewares.Auth())
	users.RegisterRoutes(api, h)

	r.Run("0.0.0.0:5000")
}
