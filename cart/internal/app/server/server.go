package server

import "route256/cart/internal/pkg/cart/model"

type Service interface {
	AddProduct(*model.Product, int64) error
	DeleteSKU(int64, int64) error
	ClearCart(int64) error
	GetCart(int64) (*model.CartItem, error)
}

type Server struct {
	Service Service
}

func NewServer(service Service) *Server {
	return &Server{
		Service: service,
	}
}
