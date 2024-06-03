package add_product

import (
	"fmt"
	"route256/cart/internal/model"
)

type Repository interface {
	GetProductCart(*model.Product, int64) (*model.Product, error)
	AddProductCart(*model.Product, int64) error
	CheckSKU(int64) (*model.Product, error)
}

type Handler struct {
	Repository Repository
}

func New(repository Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

// Добавляет товар в корзину.
// Сначала проверяется наличие товара в специальном сервисе.
// Затем получаем, если есть, количество товара добавленного ранее в корзину.
// Добавляет к нему новый объем и сохраняет в корзину
func (h *Handler) AddProduct(productRequest *model.Product, userId int64) error {
	checkSKU, err := h.Repository.CheckSKU(productRequest.SKU)
	if err != nil {
		return fmt.Errorf("s.Repository.CheckSKU %w", err)
	}

	if checkSKU.Price > 0 {
		currentProduct, err := h.Repository.GetProductCart(productRequest, userId)
		if err != nil {
			return fmt.Errorf("s.Repository.GetProductCart %w", err)
		}

		productRequest.Count += currentProduct.Count

		err = h.Repository.AddProductCart(productRequest, userId)
		if err != nil {
			return fmt.Errorf("s.Repository.AddProductCart %w", err)
		}

		return nil
	}

	if checkSKU.Price == 0 {
		return model.ErrNoProductInStock
	}

	return nil
}
