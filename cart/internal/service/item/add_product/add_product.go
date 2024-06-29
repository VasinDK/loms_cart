package add_product

import (
	"context"
	"fmt"
	"route256/cart/internal/model"
)

type Repository interface {
	GetProductCart(context.Context, *model.Product, int64) (*model.Product, error)
	AddProductCart(context.Context, *model.Product, int64) error
	CheckSKU(context.Context, chan<- *model.Product, int64) error
	StockInfo(context.Context, int64) (int64, error)
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

// AddProduct - добавляет товар в корзину.
// Сначала проверяется наличие товара в специальном сервисе.
// Затем получаем, если есть, количество товара добавленного ранее в корзину.
// Добавляет к нему новый объем и сохраняет в корзину
func (h *Handler) AddProduct(ctx context.Context, productRequest *model.Product, userId int64) error {
	ch1 := make(chan *model.Product, 1)

	err := h.Repository.CheckSKU(ctx, ch1, productRequest.SKU)
	if err != nil {
		return fmt.Errorf("h.Repository.CheckSKU %w", err)
	}
	checkSKU := <-ch1
	close(ch1)

	if productRequest.Count < 1 {
		return fmt.Errorf("AddProduct %w", fmt.Errorf("количество меньше 1"))
	}

	var countSKU int64
	ctx = context.Background()
	if checkSKU.Price > 0 {
		countSKU, err = h.Repository.StockInfo(ctx, productRequest.SKU)
		if err != nil {
			return fmt.Errorf("h.Repository.StockInfo %w", err)
		}
	}

	if countSKU >= int64(productRequest.Count) && countSKU != 0 {
		currentProduct, err := h.Repository.GetProductCart(ctx, productRequest, userId)
		if err != nil {
			return fmt.Errorf("s.Repository.GetProductCart %w", err)
		}

		productRequest.Count += currentProduct.Count

		err = h.Repository.AddProductCart(ctx, productRequest, userId)
		if err != nil {
			return fmt.Errorf("s.Repository.AddProductCart %w", err)
		}

		return nil
	}

	if countSKU < int64(productRequest.Count) {
		return fmt.Errorf("AddProduct %w", model.ErrInsufficientStock)
	}

	if checkSKU.Price == 0 {
		return fmt.Errorf("AddProduct %w", model.ErrNoProductInStock)
	}

	return nil
}
