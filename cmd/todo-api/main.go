package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/yusufpapurcu/todo-api-example/internal/config"
	"github.com/yusufpapurcu/todo-api-example/internal/database"
	"github.com/yusufpapurcu/todo-api-example/internal/router"
	"github.com/yusufpapurcu/todo-api-example/pkg/constants"
	"github.com/yusufpapurcu/todo-api-example/pkg/logger"
)

func main() {

	// Get config
	conf := config.ParseConfig(constants.ConfigPath + constants.ConfigName)

	// Create DSN and open connection
	dsn := database.CreateDSN(conf)
	db, err := database.OpenConnection(dsn)
	if err != nil {
		logger.Fatalf("%s", err.Error())
	}

	srv := &http.Server{
		MaxHeaderBytes: 10,
		Addr:           ":" + fmt.Sprint(conf.Server.Port),
		WriteTimeout:   time.Second * time.Duration(conf.Server.Timeout), // Get them from config
		ReadTimeout:    time.Second * time.Duration(conf.Server.Timeout), // Get them from config
		Handler:        router.New(db),
	}

	logger.Infof("listening on %d", conf.Server.Port)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("failed to start server: %v", err)
	}
}
