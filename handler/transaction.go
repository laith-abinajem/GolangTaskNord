package handler

import (
	"errors"
	"strconv"
	"task/model"
	"task/pkg/metrics"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (s *Server) CreateTransaction(ctx *fiber.Ctx) error {
	tx := &model.Transaction{}
	start := time.Now()

	s.logger.Info("Received request to create a new transaction")

	// Parse request body
	if err := ctx.BodyParser(&tx); err != nil {
		s.logger.Warn("Failed to parse transaction request body", "error", err)
		return NewBadRequestResponse(ctx, "Invalid transaction request")
	}

	// Validate Tenant Exists
	tenant := &model.Tenant{}
	if err := s.db.DB.First(&tenant, tx.TenantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Warn("Tenant not found", "tenantID", tx.TenantID)
			return NewNotFoundResponse(ctx, "Tenant not found")
		}
		s.logger.Error("Error fetching tenant from database", "error", err)
		return NewInternalServerErrorResponse(ctx, "Error fetching tenant")
	}

	// Validate Branch Exists
	branch := &model.Branch{}
	if err := s.db.DB.First(&branch, tx.BranchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Warn("Branch not found", "branchID", tx.BranchID)
			return NewNotFoundResponse(ctx, "Branch not found")
		}
		s.logger.Error("Error fetching branch from database", "error", err)
		return NewInternalServerErrorResponse(ctx, "Error fetching branch")
	}

	// Validate Product Exists
	product := &model.Product{}
	if err := s.db.DB.First(&product, tx.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Warn("Product not found", "productID", tx.ProductID)
			return NewNotFoundResponse(ctx, "Product not found")
		}
		s.logger.Error("Error fetching product from database", "error", err)
		return NewInternalServerErrorResponse(ctx, "Error fetching product")
	}

	// Insert into database via Core
	transaction, err := s.core.CreateTransaction(tx)
	if err != nil {
		s.logger.Error("Failed to create transaction", "error", err)
		return NewInternalServerErrorResponse(ctx, "Failed to create transaction")
	}

	// Delete cached data
	err = s.core.DeleteCachedTotalSales(tx.TenantID, tx.ProductID)
	if err != nil {
		s.logger.Warn("Error deleting cached total sales", "error", err)
	}

	err = s.core.DeleteCachedTopProducts()
	if err != nil {
		s.logger.Warn("Error deleting cached top products", "error", err)
	}

	// Log successful transaction
	s.logger.Info("Transaction created successfully", "transactionID", transaction.ID)

	// Record metrics
	metrics.TransactionsProcessed.WithLabelValues("success").Inc()
	metrics.APIRequestDuration.WithLabelValues(ctx.Route().Path, ctx.Method()).Observe(time.Since(start).Seconds())

	// Return success response
	return NewSuccessResponse(ctx, transaction)
}
func (s *Server) GetTotalSales(ctx *fiber.Ctx) error {
	tenantIDStr := ctx.Params("tenantID")
	productIDStr := ctx.Params("productID")

	s.logger.Info("Received request to get total sales", "tenantID", tenantIDStr, "productID", productIDStr)

	tenantID, err := strconv.ParseUint(tenantIDStr, 10, 32)
	if err != nil {
		s.logger.Warn("Invalid tenant ID format", "input", tenantIDStr)
		return NewBadRequestResponse(ctx, "Invalid tenant ID")
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		s.logger.Warn("Invalid product ID format", "input", productIDStr)
		return NewBadRequestResponse(ctx, "Invalid product ID")
	}

	// Validate Tenant Exists
	tenant := &model.Tenant{}
	if err := s.db.DB.First(&tenant, tenantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Warn("Tenant not found", "tenantID", tenantID)
			return NewNotFoundResponse(ctx, "Tenant not found")
		}
		s.logger.Error("Error fetching tenant from database", "error", err)
		return NewInternalServerErrorResponse(ctx, "Error fetching tenant")
	}

	totalSales, err := s.core.GetTotalSales(uint(tenantID), uint(productID))
	if err != nil {
		s.logger.Error("Failed to fetch total sales", "error", err)
		return NewInternalServerErrorResponse(ctx, "Failed to fetch sales data")
	}

	s.logger.Info("Fetched total sales successfully", "tenantID", tenantID, "productID", productID, "totalSales", totalSales)

	return NewSuccessResponse(ctx, fiber.Map{"total_sales": totalSales})
}

func (s *Server) GetTopProducts(ctx *fiber.Ctx) error {
	s.logger.Info("Received request to get top products")

	// Fetch top products from Core
	topProducts, err := s.core.GetTopProducts()
	if err != nil {
		s.logger.Error("Failed to fetch top products", "error", err)
		return NewInternalServerErrorResponse(ctx, "Failed to fetch top products")
	}

	s.logger.Info("Fetched top products successfully", "count", len(topProducts))

	return NewSuccessResponse(ctx, topProducts)
}
