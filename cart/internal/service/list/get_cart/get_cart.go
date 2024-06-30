package get_cart

import (
	"context"
	"fmt"
	"route256/cart/internal/model"
	"route256/cart/pkg/errgroup_my"
	"sort"
	"sync"
)

type Repository interface {
	CheckSKU(context.Context, int64) (*model.Product, error)
	GetCart(context.Context, int64) (map[int64]*model.Product, error)
}

type Handler struct {
	Repository Repository
}

// New - создает и возвращает Handler
func New(repository Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

// GetCart - получает содержимое конкретной корзины
func (h *Handler) GetCart(ctx context.Context, cartId int64) (*model.Cart, error) {
	var totalPrice uint32
	cart := &model.Cart{}

	productsList, err := h.Repository.GetCart(ctx, cartId)

	if err != nil {
		return cart, fmt.Errorf("s.Repository.GetCart %w", err)
	}

	mu := &sync.RWMutex{}

	eg, ctx := errgroup_my.WithContext(ctx)
	eg.SetLimitPeriod(10, 1)

	for i := range productsList {
		eg.Go(func() error {
			prod, err := h.Repository.CheckSKU(ctx, i)
			if err != nil {
				return err
			}

			mu.Lock()
			productsList[i].Price = prod.Price
			productsList[i].Name = prod.Name
			productsList[i].SKU = prod.SKU
			mu.Unlock()

			return err
		})

	}

	if err = eg.Wait(); err != nil {
		return nil, fmt.Errorf("eg.Wait %w", err)
	}

	products := make([]*model.Product, 0)

	for _, v := range productsList {
		products = append(products, v)
		totalPrice += v.Price * uint32(v.Count)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].SKU < products[j].SKU
	})

	cart.Items = products
	cart.TotalPrice = totalPrice

	return cart, nil
}
