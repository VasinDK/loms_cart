package clear_cart

type Repository interface {
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

// ClearCart - отчищает корзину, удали ее полностью
func (h *Handler) ClearCart(cartId int64) error {
	err := h.Repository.ClearCart(cartId)
	return err
}
