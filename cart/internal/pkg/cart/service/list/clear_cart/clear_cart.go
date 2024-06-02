package clear_cart

type Repository interface {
	ClearCart(int64) error
}

type Handler struct {
	Repository Repository
}

func New(repository Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

// Отчищает корзину, удали ее полностью
func (h *Handler) ClearCart(cartId int64) error {
	err := h.Repository.ClearCart(cartId)
	return err
}
