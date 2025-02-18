package app

import (
	"task/model"
	"task/pkg/config"
	"task/pkg/db"
)

func Migrate(cfg *config.Config, db *db.MySQLDB) error {
	models := []interface{}{
		&model.Transaction{},
		&model.Branch{},
		&model.Tenant{},
		&model.Product{},
	}

	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}

	return nil
}
