package main

import (
	"net/http"
	"os"
	"time"

	"github.com/yusufpapurcu/todo-api-example/internal/database"
	"github.com/yusufpapurcu/todo-api-example/internal/router"
	"github.com/yusufpapurcu/todo-api-example/pkg/logger"
)

func main() {

	// TODO: Create config package for here
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err := database.OpenConnection(dsn)
	if err != nil {
		logger.Fatalf("%s", err.Error())
	}

	srv := &http.Server{
		MaxHeaderBytes: 10, // 10 MB
		Addr:           ":" + os.Getenv("SERVER_PORT"),
		WriteTimeout:   time.Second * time.Duration(20), // Get them from config
		ReadTimeout:    time.Second * time.Duration(20), // Get them from config
		IdleTimeout:    time.Second * 60,
		Handler:        router.New(db),
	}

	logger.Infof("listening on %s", os.Getenv("SERVER_PORT"))
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}
