package server

import (
	"route256/cart/internal/pkg/cart/model"

	"github.com/go-playground/validator/v10"
)

// Корзина ответ
type CartResponse struct {
	Items      []*ProductResponse `json:"items"`
	TotalPrice uint32             `json:"total_price"`
}

// Товар ответ
type ProductResponse struct {
	SKU   int64  `json:"sku_id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
	Count uint16 `json:"count"`
}

// Товар запрос
type ProductRequest struct {
	Count uint16 `json:"count" validate:"gte=1"`
}

type Service interface {
	AddProduct(*model.Product, int64) error
	DeleteProductCart(int64, int64) error
	ClearCart(int64) error
	GetCart(int64) (*model.Cart, error)
}

type Server struct {
	Service Service
}

// Инициализирует экземпляр транспорта
func NewServer(service Service) *Server {
	return &Server{
		Service: service,
	}
}

var validate = validator.New(validator.WithRequiredStructEnabled())
