package checkout

import (
	"fmt"
	"route256/cart/internal/model"
	"route256/cart/internal/service/list/clear_cart"
	"route256/cart/internal/service/list/get_cart"
)

type Repository interface {
	Checkout(int64, []*model.Product) (int64, error)
	CheckSKU(int64) (*model.Product, error)
	GetCart(int64) (map[int64]*model.Product, error)
	ClearCart(int64) error
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

// Checkout - создает ордер в order storage``
func (h *Handler) Checkout(userId int64) (int64, error) {
	cart, err := get_cart.New(h.Repository).GetCart(userId)
	if err != nil {
		return 0, fmt.Errorf("get_cart.New().GetCart() %w", err)
	}
	orderId, err := h.Repository.Checkout(userId, cart.Items)
	if err != nil {
		return 0, fmt.Errorf("h.Repository.Checkout %w", err)
	}

	err = clear_cart.New(h.Repository).ClearCart(userId)
	if err != nil {
		return 0, fmt.Errorf("clear_cart.New().ClearCart() %w", err)
	}

	return orderId, nil
}