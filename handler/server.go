package handler

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"task/internal/core"
	"task/pkg/cache"
	"task/pkg/config"
	"task/pkg/db"
	"task/pkg/logger"
	"task/pkg/metrics"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	config    *config.Config
	logger    *logger.Logger
	app       *fiber.App
	db        *db.MySQLDB
	cache     *cache.Cache
	validator *validator.Validate

	// Services
	core *core.Core
}

func NewServer(cfg *config.Config, logger *logger.Logger, db *db.MySQLDB, cache *cache.Cache) (*Server, error) {
	v := validator.New()
	c, err := core.NewCore(cfg, logger, db, cache, v)
	if err != nil {
		return nil, fmt.Errorf("failed to create core: %w", err)
	}

	// Register Prometheus metrics
	metrics.RegisterMetrics()

	server := &Server{
		config:    cfg,
		logger:    logger,
		app:       fiber.New(),
		db:        db,
		cache:     cache,
		validator: v,
		core:      c,
	}

	// Expose Prometheus metrics at `/metrics`
	server.app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	return server, nil
}
func (s *Server) Start() {
	s.logger.Info("Starting server...")
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))
	// Register routes
	s.RegisterRoutes()

	// Create a channel to listen for OS interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a separate goroutine
	go func() {
		if err := s.app.Listen(":" + s.config.Server.Port); err != nil {
			s.logger.Fatal("Failed to start server:", err)
		}
	}()

	// Wait for an interrupt signal (Ctrl+C)
	<-stop

	s.logger.Info("Shutting down gracefully...")

	// Gracefully shutdown Fiber
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.app.Shutdown(); err != nil {
		s.logger.Error("Error shutting down server:", err)
	}

	s.logger.Info("Server stopped.")
}
