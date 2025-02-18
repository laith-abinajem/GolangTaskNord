package core

import (
	"task/model"
	"task/pkg/metrics"
)

func (c *Core) CreateTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	transaction.TotalAmount = float64(transaction.Quantity) * transaction.PricePerUnit

	c.transactionQueue <- transaction

	return transaction, nil
}
func (c *Core) GetTotalSales(tenantID, productID uint) (float64, error) {
	var totalSales float64

	CachedTotalSales, err := c.GetCachedTotalSales(tenantID, productID)
	if err != nil {
		c.logger.Warn("Cache fetch error:", err)
	}

	if CachedTotalSales > 0 {
		c.logger.Info("Returning cached total sales:", CachedTotalSales)
		return CachedTotalSales, nil
	}
	metrics.CacheMisses.Inc()

	err = c.db.DB.Model(&model.Transaction{}).
		Where("tenant_id = ? AND product_id = ?", tenantID, productID).
		Select("SUM(total_amount)").
		Scan(&totalSales).Error
	if err != nil {
		c.logger.Error("DB query failed:", err)
		return 0, err
	}

	err = c.CacheTotalSales(tenantID, productID, totalSales)
	if err != nil {
		c.logger.Warn("Failed to cache total sales:", err)
	}

	return totalSales, nil
}

func (c *Core) GetTopProducts() ([]*model.TopProduct, error) {
	topProducts := []*model.TopProduct{}

	// Check cache first
	CachedTopProducts, err := c.GetCachedTopProducts()
	if err != nil {
		c.logger.Warn("Cache fetch error:", err)
	}

	// If cache is valid, return cached data
	if len(CachedTopProducts) > 0 {
		c.logger.Info("Returning cached top products")
		return CachedTopProducts, nil
	}

	// Fetch from DB
	err = c.db.DB.
		Model(&model.Transaction{}).
		Select("Product.id AS product_id, Product.name, Product.price, SUM(Transaction.quantity) AS total_quantity").
		Joins("JOIN Product ON Transaction.product_id = Product.id").
		Group("Product.id, Product.name, Product.price").
		Order("total_quantity DESC").
		Scan(&topProducts).Error
	if err != nil {
		c.logger.Error("Failed to fetch top products from DB:", err)
		return nil, err
	}

	err = c.CacheTopProducts(topProducts)
	if err != nil {
		c.logger.Warn("Failed to cache top products:", err)
	}

	return topProducts, nil
}
