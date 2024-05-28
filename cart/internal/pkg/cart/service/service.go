package service

import (
	"fmt"
	"route256/cart/internal/pkg/cart/model"
)

type Service struct {
	Repository Repository
}

type ProductAdded struct {
	SKU    int64
	UserId int64
	Count  uint16
}

type Repository interface {
	AddProductCart(*model.Product, int64) error
	CheckSKU(int64) (bool, error)
	DeleteSKU(int64, int64) error
	ClearCart(int64) error
	GetCart(int64) (*model.CartItem, error)
}

func NewService(repo Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) AddProduct(productRequest *model.Product, userId int64) error {
	checkSKU, err := s.Repository.CheckSKU(productRequest.SKU)
	if err != nil {
		return fmt.Errorf("s.Repository.CheckSKU %w", err)
	}

	if checkSKU {
		err := s.Repository.AddProductCart(productRequest, userId)
		if err != nil {
			return fmt.Errorf("s.Repository.AddProductCart %w", err)
		}

		return nil
	}

	if !checkSKU {
		return ErrNoProductInStock
	}

	return nil
}

func (s *Service) DeleteSKU(cartId, sku int64) error {
	err := s.Repository.DeleteSKU(cartId, sku)
	return err
}

func (s *Service) ClearCart(cartId int64) error {
	err := s.Repository.ClearCart(cartId)
	return err
}

func (s *Service) GetCart(cartId int64) (*model.CartItem, error) {
	return s.Repository.GetCart(cartId)
}
