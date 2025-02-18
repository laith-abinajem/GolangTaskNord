package db

import (
	"fmt"
	"task/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// MySQLDB represents the database instance
type MySQLDB struct {
	*gorm.DB
	config *config.Config
}

// NewMySQLDB initializes a MySQL database connection
func NewMySQLDB(cfg *config.Config) (*MySQLDB, error) {
	fmt.Println("Connecting to MySQL with:")
	fmt.Printf("Host: %s, Port: %s, User: %s, Database: %s\n",
		cfg.MySQL.MysqlHost, cfg.MySQL.MysqlPort, cfg.MySQL.MysqlUser, cfg.MySQL.MYSQLDB)

	if cfg.MySQL.MysqlHost == "" || cfg.MySQL.MysqlPort == "" {
		return nil, fmt.Errorf("  MySQL host or port is missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.MysqlUser,
		cfg.MySQL.MysqlPassword,
		cfg.MySQL.MysqlHost,
		cfg.MySQL.MysqlPort,
		cfg.MySQL.MYSQLDB,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		fmt.Println("  Database Connection Failed:", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println(" MySQL Connection Established Successfully")
	return &MySQLDB{db, cfg}, nil
}
