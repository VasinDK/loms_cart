package get_cart

import (
	"fmt"
	"route256/cart/internal/model"
	"sort"
	"strings"
)

type Repository interface {
	CheckSKU(int64) (*model.Product, error)
	GetCart(int64) (map[int64]*model.Product, error)
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
func (h *Handler) GetCart(cartId int64) (*model.Cart, error) {
	var totalPrice uint32
	cart := &model.Cart{}

	productsList, err := h.Repository.GetCart(cartId)
	if err != nil {
		return cart, fmt.Errorf("s.Repository.GetCart %w", err)
	}

	products := make([]*model.Product, 0, len(productsList))

	errorsSKU := make([]string, 0)

	for i := range productsList {
		item, err := h.Repository.CheckSKU(i)

		if err != nil {
			errorsSKU = append(errorsSKU, err.Error())
			continue
		}

		item.SKU = i
		item.Count = productsList[i].Count

		products = append(products, item)

		totalPrice += item.Price * uint32(productsList[i].Count)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].SKU < products[j].SKU
	})

	cart.Items = products
	cart.TotalPrice = totalPrice

	if len(errorsSKU) > 0 {
		return cart, fmt.Errorf("range productsList s.Repository.CheckSKU Errors: %v",
			strings.Join(errorsSKU, ", "))
	}

	return cart, nil
}
