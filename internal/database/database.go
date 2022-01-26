package database

import (
	"fmt"

	"github.com/yusufpapurcu/todo-api-example/internal/config"
	"github.com/yusufpapurcu/todo-api-example/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.Task{})
	return db, err
}

func CreateDSN(conf config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.DBName,
		conf.Database.Port,
		conf.Database.TimeZone)
}
