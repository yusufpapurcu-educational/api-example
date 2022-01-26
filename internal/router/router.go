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
	handler.GET("/task", api.GetAllTasks(db))
	handler.POST("/task", api.CreateTask(db))
	handler.GET("/task/:id", api.GetTask(db))
	handler.PUT("/task/:id", api.UpdateTask(db))
	handler.DELETE("/task/:id", api.DeleteTask(db))
}
