package service

import (
	"fmt"
	"route256/cart/internal/pkg/cart/model"
	"sort"
	"strings"
)

// Структура сервис
type Service struct {
	Repository Repository
}

type Repository interface {
	GetProductCart(*model.Product, int64) (*model.Product, error)
	AddProductCart(*model.Product, int64) error
	DeleteProductCart(int64, int64) error
	CheckSKU(int64) (*model.Product, error)
	ClearCart(int64) error
	GetCart(int64) (map[int64]*model.Product, error)
}

// Инициализирует сервис
func NewService(repo Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

// Добавляет товар в корзину.
// Сначала проверяется наличие товара в специальном сервисе.
// Затем получаем, если есть, количество товара добавленного ранее в корзину.
// Добавляет к нему новый объем и сохраняет в корзину
func (s *Service) AddProduct(productRequest *model.Product, userId int64) error {
	checkSKU, err := s.Repository.CheckSKU(productRequest.SKU)
	if err != nil {
		return fmt.Errorf("s.Repository.CheckSKU %w", err)
	}

	if checkSKU.Price > 0 {
		currentProduct, err := s.Repository.GetProductCart(productRequest, userId)
		if err != nil {
			return fmt.Errorf("s.Repository.GetProductCart %w", err)
		}

		productRequest.Count += currentProduct.Count

		err = s.Repository.AddProductCart(productRequest, userId)
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

// Удаляет товар из корзины
func (s *Service) DeleteProductCart(cartId, sku int64) error {
	err := s.Repository.DeleteProductCart(cartId, sku)
	return err
}

// Отчищает корзину, удали ее полностью
func (s *Service) ClearCart(cartId int64) error {
	err := s.Repository.ClearCart(cartId)
	return err
}

// Получает содержимое конкретной корзины
func (s *Service) GetCart(cartId int64) (*model.Cart, error) {
	var totalPrice uint32
	cart := &model.Cart{}

	productsList, err := s.Repository.GetCart(cartId)
	if err != nil {
		return cart, fmt.Errorf("s.Repository.GetCart %w", err)
	}

	products := make([]*model.Product, 0, len(productsList))

	errorsSKU := make([]string, 0)

	for i := range productsList {
		item, err := s.Repository.CheckSKU(i)
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
