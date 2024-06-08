package http_handlers

import (
	"github.com/go-playground/validator/v10"
)

// CartResponse - корзина, ответ
type CartResponse struct {
	Items      []*ProductResponse `json:"items"`
	TotalPrice uint32             `json:"total_price"`
}

// ProductResponse - товар, ответ
type ProductResponse struct {
	SKU   int64  `json:"sku_id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
	Count uint16 `json:"count"`
}

// ProductRequest - товар, запрос
type ProductRequest struct {
	Count uint16 `json:"count" validate:"gte=1"`
}

type Server struct{}

// New - инициализирует экземпляр транспорта
func New() *Server {
	return &Server{}
}

var validate = validator.New(validator.WithRequiredStructEnabled())
