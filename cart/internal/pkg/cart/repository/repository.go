package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"route256/cart/internal/pkg/cart/model"
)

type Carts map[int64]map[int64]*model.CartItem

type Repository struct {
	Carts Carts
}

func NewRepository() *Repository {
	return &Repository{
		Carts: make(Carts),
	}
}

func (r *Repository) CheckSKU(sku int64) (bool, error) {
	type CheckSkuRequest struct {
		Token string `json:"token"`
		Sku   int64  `json:"sku"`
	}

	type CheckSkuResponse struct {
		Code      int     `json:"code"`
		Name      string  `json:"name"`
		Price     float32 `json:"price"`
		ErrorMess string  `json:"message"`
	}

	bodyCheckSKU := CheckSkuRequest{
		Token: "testtoken",
		Sku:   sku,
	}

	jsonBodyCheckSKU, err := json.Marshal(bodyCheckSKU)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest(
		"POST",
		"http://route256.pavl.uk:8080/get_product",
		bytes.NewReader(jsonBodyCheckSKU),
	)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	SkuResponse := CheckSkuResponse{}
	json.NewDecoder(resp.Body).Decode(&SkuResponse)

	if SkuResponse.Price > 0 {
		return true, nil
	}

	return false, nil
}

// TODO Разложить. Логику вынести в сервис
func (r *Repository) AddProductCart(productRequest *model.Product, cartId int64) error {
	if _, ok := r.Carts[cartId]; !ok {
		r.Carts[cartId] = make(map[int64]*model.CartItem)
	}

	var ctn uint16
	if basket, ok := r.Carts[cartId][productRequest.SKU]; ok {
		ctn = basket.Count
	}

	r.Carts[cartId][productRequest.SKU] = &model.CartItem{
		Count: productRequest.Count + ctn,
	}

	return nil
}

func (r *Repository) DeleteSKU(cartId, sku int64) error {
	delete(r.Carts[cartId], sku)
	return nil
}

func (r *Repository) ClearCart(cartId int64) error {
	delete(r.Carts, cartId)
	return nil
}

func (r *Repository) GetCart(cartId int64) (map[int64]*model.CartItem, error) {
	return r.Carts[cartId], nil
}
