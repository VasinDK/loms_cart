package delete_item

type Repository interface {
	DeleteProductCart(int64, int64) error
}

type Handler struct {
	Repository Repository
}

func New(repository Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

// Удаляет товар из корзины
func (h *Handler) DeleteProductCart(cartId, sku int64) error {
	err := h.Repository.DeleteProductCart(cartId, sku)
	return err
}
