package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTasks(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {}
}

func CreateTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {}
}

func GetTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {}
}

func UpdateTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {}
}

func DeleteTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {}
}
