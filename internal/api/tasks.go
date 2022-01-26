package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yusufpapurcu/todo-api-example/model"
	"gorm.io/gorm"
)

func GetAllTasks(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var all_tasks []model.Task
		db.Find(&all_tasks)
		c.JSON(200, all_tasks)
	}
}

func CreateTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var wanted model.Task
		c.Bind(&wanted)
		validator.New().Struct(wanted)
		res := db.Create(&wanted)
		if res.Error != nil {
			c.JSON(400, map[string]string{
				"error": res.Error.Error(),
			})
		} else {
			c.JSON(200, map[string]int64{
				"rows_affected": res.RowsAffected,
			})
		}
	}
}

func GetTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var wanted model.Task
		if res := db.First(&wanted, id); res.Error != nil {
			c.JSON(400, map[string]string{
				"error": res.Error.Error(),
			})
		} else {
			c.JSON(200, wanted)
		}
	}
}

func UpdateTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var from model.Task
		db.First(&from, id)

		var wanted model.Task
		c.Bind(&wanted)

		wanted.ID = from.ID
		if res := db.Save(&wanted); res.Error != nil {
			c.JSON(400, map[string]string{
				"error": res.Error.Error(),
			})
		} else {
			c.JSON(200, map[string]int64{
				"rows_affected": int64(res.RowsAffected),
			})
		}
	}
}

func DeleteTask(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		var from model.Task
		db.First(&from, id)
		if res := db.Delete(&from); res.Error != nil {
			c.JSON(400, map[string]string{
				"error": res.Error.Error(),
			})
		} else {
			c.JSON(200, map[string]int64{
				"rows_affected": res.RowsAffected,
			})
		}
	}
}
