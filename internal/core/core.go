package core

import (
	"errors"
	"sync"
	"task/model"
	"task/pkg/cache"
	"task/pkg/config"
	"task/pkg/db"
	"task/pkg/logger"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Core struct {
	config *config.Config
	logger *logger.Logger
	db     *db.MySQLDB
	cache  *cache.Cache

	validator *validator.Validate

	transactionQueue chan *model.Transaction
	workerWg         sync.WaitGroup
	mu               sync.Mutex
}

func NewCore(c *config.Config, l *logger.Logger, db *db.MySQLDB, cache *cache.Cache, v *validator.Validate) (*Core, error) {
	core := &Core{
		config:           c,
		logger:           l,
		db:               db,
		cache:            cache,
		validator:        v,
		transactionQueue: make(chan *model.Transaction, 1000),
	}

	core.StartWorkers(5)

	return core, nil
}
func (c *Core) StartWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		c.workerWg.Add(1)
		go c.worker()
	}
}
func (c *Core) worker() {
	defer c.workerWg.Done()

	for transaction := range c.transactionQueue {
		c.processTransaction(transaction)
	}
}
func (c *Core) processTransaction(transaction *model.Transaction) {
	c.mu.Lock()
	defer c.mu.Unlock()

	maxRetries := 3
	delay := time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := c.db.DB.Create(transaction).Error
		if err == nil {
			break
		}

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.logger.Error("Duplicate transaction detected, skipping.")
			return
		}

		c.logger.Warn("Transaction failed, retrying...", "attempt", attempt, "error", err)
		time.Sleep(delay)
	}

	if err := c.CacheTransaction(transaction); err != nil {
		c.logger.Error("Failed to cache transaction", "error", err)
	}
}
