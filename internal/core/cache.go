package core

import (
	"context"
	"encoding/json"
	"task/internal/keys"
	"task/model"

	"github.com/redis/go-redis/v9"
)

const (
	OneDay                 = 60 * 60 * 24
	OneWeek                = OneDay * 7
	_TRANSACTION_CACHE_TTL = OneDay
	_PROD_TOP_CACHE_TTL    = OneDay
	_SALES_CACHE_TTL       = OneDay
)

func (c *Core) CacheTransaction(transaction *model.Transaction) error {
	marshaled, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	return c.cache.Set(context.TODO(), keys.KEY_Transaction(transaction.ID), marshaled, _TRANSACTION_CACHE_TTL)

}

func (c *Core) CacheTotalSales(tenantID, productID uint, totalSales float64) error {
	marshaled, err := json.Marshal(totalSales)
	if err != nil {
		return err
	}
	return c.cache.Set(context.TODO(), keys.KEY_TotalSales(tenantID, productID), marshaled, _SALES_CACHE_TTL)
}

func (c *Core) GetCachedTotalSales(tenantID, productID uint) (float64, error) {
	var totalSales float64

	// Fetch data from cache
	cachedData, err := c.cache.Get(context.TODO(), keys.KEY_TotalSales(tenantID, productID))
	if err != nil {
		if err != redis.Nil {
			c.logger.Warn("Cache error:", err)
			return 0, err
		}
		return 0, nil
	}

	err = json.Unmarshal([]byte(cachedData), &totalSales)
	if err != nil {
		c.logger.Error("Failed to unmarshal cached sales data:", err)
		return 0, err
	}

	return totalSales, nil
}
func (c *Core) DeleteCachedTotalSales(tenantID, productID uint) error {
	return c.cache.Delete(context.TODO(), keys.KEY_TotalSales(tenantID, productID))
}

func (c *Core) CacheTopProducts(topProducts []*model.TopProduct) error {

	marshaled, err := json.Marshal(topProducts)
	if err != nil {
		return err
	}
	return c.cache.Set(context.TODO(), keys.KEY_TopProducts(), marshaled, _PROD_TOP_CACHE_TTL)

}

func (c *Core) GetCachedTopProducts() ([]*model.TopProduct, error) {
	topProducts := []*model.TopProduct{}

	cachedData, err := c.cache.Get(context.TODO(), keys.KEY_TopProducts())
	if err != nil {
		if err != redis.Nil {
			c.logger.Warn("Cache fetch error for top products:", err)
		}
		return nil, nil
	}

	err = json.Unmarshal([]byte(cachedData), &topProducts)
	if err != nil {
		c.logger.Error("Failed to unmarshal cached top products:", err)
		return nil, err
	}

	return topProducts, nil
}
func (c *Core) DeleteCachedTopProducts() error {
	return c.cache.Delete(context.TODO(), keys.KEY_TopProducts())
}
