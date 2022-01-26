package database

import (
	"github.com/yusufpapurcu/todo-api-example/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.Task{})
	return db, err
}
