package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func InitializeExternalRoutes (router *gin.Engine, handler *Handler) {
	v1 :=router.Group("v1") 
	{
		initializeRoutes(v1, handler)
	}
	swagger := router.Group("swagger")
	{
		initializeDocsRoutes(swagger)
	}
}

func initializeRoutes(parent *gin.RouterGroup, handler *Handler) {
	parent.Post("create", handler.CreateTask)
	parent.POST("/tasks", handler.CreateTask)
	parent.GET("/tasks/:id", handler.GetTask)
	parent.GET("/tasks", handler.GetTasks)
	parent.PUT("/tasks/:id", handler.UpdateTask)
	parent.DELETE("/tasks/:id", handler.DeleteTask)
	parent.GET("/tasks/pending", handler.GetPendingTasks)
	parent.GET("/tasks/completed", handler.GetCompletedTasks)
	parent.GET("/tasks/date", handler.GetTasksByDate)
}

func initializeDocsRoutes(parent *gin.RouterGroup) {
	parent.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}