package server

import (
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

type Server struct{}

// Инициализирует экземпляр транспорта
func NewServer() *Server {
	return &Server{}
}

var validate = validator.New(validator.WithRequiredStructEnabled())
