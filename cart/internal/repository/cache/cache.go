package cache

import (
	"context"
	"fmt"
	"route256/cart/internal/model"
	"route256/cart/internal/pkg/ring_cache"
	"route256/cart/internal/repository"
	"strconv"
	"sync"
)

// CacheRepo - репозиторий который кэширует значения того или иного метода
type CacheRepo struct {
	*repository.Repository
	checkSKUCache ring_cache.Cacher
	//... + кеш других методов при необходимости
}

type Config interface {
	GetSizeBufferCache() int64
}

// New - возвращает новый CacheRepo
func New(config Config, repo *repository.Repository) *CacheRepo {
	return &CacheRepo{
		Repository: repo,
		checkSKUCache: ring_cache.Cacher{
			SizeBuffer: config.GetSizeBufferCache(),
			Cache:      sync.Map{},
			RingBuffer: make([]string, config.GetSizeBufferCache()),
		},
	}
}

// CheckSKU - метод значения которого кэшируются и соответственно обрабатываются
func (c *CacheRepo) CheckSKU(ctx context.Context, sku int64) (*model.Product, error) {
	skuStr := strconv.FormatInt(sku, 10)

	if v, ok := c.checkSKUCache.Get(ctx, skuStr); ok {
		product := (v).(*model.Product)

		return product, nil
	}

	product, err := c.Repository.CheckSKU(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("c.Repository.CheckSKU %w", err)
	}

	c.checkSKUCache.Set(ctx, skuStr, product)

	return product, nil
}
