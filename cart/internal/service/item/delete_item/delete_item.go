package delete_item

import "context"

type Repository interface {
	DeleteProductCart(context.Context, int64, int64) error
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

// DeleteProductCart - удаляет товар из корзины
func (h *Handler) DeleteProductCart(ctx context.Context, cartId, sku int64) error {
	err := h.Repository.DeleteProductCart(ctx, cartId, sku)
	return err
}
