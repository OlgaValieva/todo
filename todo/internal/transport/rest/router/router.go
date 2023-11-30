package router

import (
	"todo/config"
	domain "todo/internal/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(cfg config.AppConfig) *gin.Engine {
	router := gin.Default()

	if cfg.Debug {
		router.Use(cors.Default())
	}

	router.NoRoute (func(c *gin.Context){
		c.JSON(404, gin.H{"code": domain.ErrPageNotFound, "message": "Page not found"})
	})
	return router
}