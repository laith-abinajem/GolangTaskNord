package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

type Logger struct {
	LogFile string `envconfig:"LOG_FILE" validate:"required"`
}

type Server struct {
	Port string `envconfig:"SERVER_PORT" validate:"required,numeric"`
}

type Debug struct {
	Debug bool `envconfig:"DEBUG"`
}

type MySQL struct {
	MysqlHost     string `envconfig:"MYSQL_HOST" validate:"required"`
	MysqlPort     string `envconfig:"MYSQL_PORT" validate:"required"`
	MysqlUser     string `envconfig:"MYSQL_USER" validate:"required"`
	MysqlPassword string `envconfig:"MYSQL_PASSWORD" validate:"required"`
	MYSQLDB       string `envconfig:"MYSQL_DB" validate:"required"`
}

type Redis struct {
	Url string `envconfig:"REDIS_URL" validate:"required"`
}
type Config struct {
	MySQL  MySQL
	Logger Logger
	Server Server
	Debug  Debug
	Redis  Redis
}

func NewConfig() (*Config, error) {
	fmt.Println("Loading configuration...")

	debug := os.Getenv("DEBUG")
	fmt.Printf("DEBUG: %s\n", debug)

	cfg := &Config{}

	// Load environment variables into the Config struct using envconfig
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Validate the config struct
	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	// Print out the loaded config (for testing purposes)
	log.Printf("Configuration Loaded: %+v\n\n", cfg)
	return cfg, nil
}

func NewTestConfig() (*Config, error) {
	cfg := &Config{
		Server: Server{
			Port: "8080",
		},
	}

	return cfg, nil
}
