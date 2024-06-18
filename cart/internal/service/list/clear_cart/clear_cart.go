package clear_cart

import "context"

type Repository interface {
	ClearCart(context.Context, int64) error
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

// ClearCart - отчищает корзину, удали ее полностью
func (h *Handler) ClearCart(ctx context.Context, cartId int64) error {
	err := h.Repository.ClearCart(ctx, cartId)
	return err
}
