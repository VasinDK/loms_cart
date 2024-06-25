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
	CheckSKU(context.Context, chan<- *model.Product, int64) error
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

	products := make([]*model.Product, 0, len(productsList))

	eg, ctx := errgroup_my.WithContext(ctx)
	eg.SetLimitPeriod(10, 1)

	ch1 := make(chan *model.Product)

	for i := range productsList {
		eg.Go(func() error {
			err := h.Repository.CheckSKU(ctx, ch1, i)
			return err
		})

	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for j := range ch1 {
			item := j
			item.Count = productsList[j.SKU].Count
			products = append(products, item)

			totalPrice += item.Price * uint32(productsList[j.SKU].Count)
		}
		wg.Done()
	}()

	err = eg.Wait()
	if err != nil {
		return nil, fmt.Errorf("eg.Wait %w", err)
	}

	close(ch1)

	wg.Wait()

	sort.Slice(products, func(i, j int) bool {
		return products[i].SKU < products[j].SKU
	})

	cart.Items = products
	cart.TotalPrice = totalPrice

	return cart, nil
}
