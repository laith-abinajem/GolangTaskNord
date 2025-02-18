package main

import (
	"fmt"
	"task/app"
	"task/pkg/config"
	"task/pkg/db"
	"task/pkg/logger"
)

func main() {

	cfg, _ := config.NewConfig()
	log, err := logger.NewLogger(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %s", err))
	}

	// DB connection
	log.Info("Creating db connection...")
	db, err := db.NewMySQLDB(cfg)
	if err != nil {
		log.Fatalf("failed to create db connection: %s", err)
	}

	err = app.Migrate(cfg, db)
	if err != nil {
		log.Fatalf("failed to migrate: %s", err)
	}

}
