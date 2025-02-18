package app

import (
	"fmt"
	"task/handler"
	"task/pkg/cache"
	"task/pkg/config"
	"task/pkg/db"
	"task/pkg/logger"
)

func Start() {

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("  Failed to load config:", err)
		panic(err)
	}
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
	log.Info("DB connection successful")

	// Migrate
	log.Info("Migrating...")
	err = Migrate(cfg, db)
	if err != nil {
		log.Fatalf("failed to migrate: %s", err)
	}
	log.Info("Migration successful")

	cache, err := cache.NewCache(cfg)
	if err != nil {
		log.Fatalf("failed to create cache: %s", err)
	}

	// Create server
	s, err := handler.NewServer(cfg, log, db, cache)
	if err != nil {
		log.Fatalf("failed to create server: %s", err)
	}

	log.Info("Creating server...")
	s.Start()
}
