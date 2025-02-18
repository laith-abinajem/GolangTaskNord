package core

import (
	"fmt"
	"task/model"

	"gorm.io/gorm"
)

func (c *Core) SeedBranches(db *gorm.DB) error {
	branches := []model.Branch{
		{Name: "Downtown Branch", TenantID: 1},
		{Name: "Industrial Zone Branch", TenantID: 1},
	}

	for _, branch := range branches {
		if err := db.Create(&branch).Error; err != nil {
			return fmt.Errorf("failed to seed branches: %w", err)
		}
	}

	fmt.Println(" Branches seeded successfully")
	return nil
}

func (c *Core) SeedProducts(db *gorm.DB) error {
	products := []model.Product{
		{Name: "Laptop", Price: 1500.0},
		{Name: "Smartphone", Price: 800.0},
		{Name: "Tablet", Price: 500.0},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			return fmt.Errorf("failed to seed products: %w", err)
		}
	}

	fmt.Println(" Products seeded successfully")
	return nil
}

func (c *Core) SeedTenants(db *gorm.DB) error {
	tenants := []model.Tenant{
		{Name: "Tech Corp"},
		{Name: "Retail Empire"},
	}

	for _, tenant := range tenants {
		if err := db.Create(&tenant).Error; err != nil {
			return fmt.Errorf("failed to seed tenants: %w", err)
		}
	}

	fmt.Println(" Tenants seeded successfully")
	return nil
}

func (c *Core) SeedDatabase(db *gorm.DB) error {
	if err := c.SeedTenants(db); err != nil {
		return err
	}
	if err := c.SeedBranches(db); err != nil {
		return err
	}
	if err := c.SeedProducts(db); err != nil {
		return err
	}

	fmt.Println("🎉 Database seeding completed successfully")
	return nil
}
