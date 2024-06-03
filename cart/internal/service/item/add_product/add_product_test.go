package add_product

import (
	"testing"

	"route256/cart/internal/service/item/add_product/mock"

	"github.com/gojuno/minimock"
)

func TestAddProduct(t *testing.T) {
	ctrl := minimock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock.NewRepositoryMock()

}
