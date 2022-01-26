package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yusufpapurcu/todo-api-example/model"
	"gorm.io/gorm"
)

func GetAllTasks(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var all_tasks []model.Task
		db.Find(&all_tasks)
		res, err := json.Marshal(all_tasks)
		if err != nil {
			c.AbortWithError(400, err)
		}
		c.Writer.Write(res)
	}
}

func CreateTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var wanted model.Task
		c.Bind(&wanted)
		validator.New().Struct(wanted)
		res := db.Create(&wanted)

		fmt.Fprintf(c.Writer, "{\nerror:%v,\nrows_affected:%d\n}", res.Error, res.RowsAffected)
	}
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
