package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusufpapurcu/todo-api-example/internal/api"
	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	handler := gin.Default()

	setRoutes(handler, db)
	return handler
}

func setRoutes(handler *gin.Engine, db *gorm.DB) {
	task_handlers := handler.Group("/task")
	{
		task_handlers.GET("/", api.GetAllTasks(db))
		task_handlers.POST("/", api.CreateTask(db))
		task_handlers.GET("/:id", api.GetTask(db))
		task_handlers.PUT("/:id", api.UpdateTask(db))
		task_handlers.DELETE("/:id", api.DeleteTask(db))
	}
}
