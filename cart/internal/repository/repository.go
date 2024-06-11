package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"route256/cart/internal/model"
)

type Config interface {
	GetTokenStore() string
	GetAddressStore() string
}

// Carts - структура: "Корзины". id пользователя есть id конкретной корзины.
// Id корзины соответствует первому ключу. SKU товара соответствует
// второму ключу. Значение - продукт
type Carts map[int64]map[int64]*model.Product

// Repository - структура репозитория
type Repository struct {
	Carts  Carts
	Config Config
}

// NewRepository - инициализирует репозиторий
func NewRepository(config Config) *Repository {
	return &Repository{
		Carts:  make(Carts),
		Config: config,
	}
}

// CheckSKU - проверяет наличие sku на удаленном сервере
func (r *Repository) CheckSKU(sku int64) (*model.Product, error) {
	// CheckSkuRequest - структура для отправки запроса SKU
	type CheckSkuRequest struct {
		Token string `json:"token"`
		Sku   int64  `json:"sku"`
	}

	// CheckSkuResponse - структура для получения ответа о наличии SKU
	type CheckSkuResponse struct {
		Code      int64  `json:"code"`
		Name      string `json:"name"`
		Price     uint32 `json:"price"`
		ErrorMess string `json:"message"`
	}

	bodyCheckSKU := CheckSkuRequest{
		Token: r.Config.GetTokenStore(),
		Sku:   sku,
	}

	responseSKU := &model.Product{}

	jsonBodyCheckSKU, err := json.Marshal(bodyCheckSKU)
	if err != nil {
		return responseSKU, err
	}

	req, err := http.NewRequest(
		"POST",
		r.Config.GetAddressStore(),
		bytes.NewReader(jsonBodyCheckSKU),
	)
	if err != nil {
		return responseSKU, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return responseSKU, err
	}
	defer resp.Body.Close()

	SkuResponseCheck := CheckSkuResponse{}
	json.NewDecoder(resp.Body).Decode(&SkuResponseCheck)

	// Если в ответе есть цена SKU считаем что товар доступен
	if SkuResponseCheck.Price > 0 {
		responseSKU.Name = SkuResponseCheck.Name
		responseSKU.Price = SkuResponseCheck.Price
		return responseSKU, nil
	}

	return responseSKU, model.ErrNoProductInStock
}

// GetProductCart - получает конкретный товар из корзины пользователя
func (r *Repository) GetProductCart(productRequest *model.Product, cartId int64) (*model.Product, error) {
	item := &model.Product{}

	if _, ok := r.Carts[cartId]; ok {
		if v, ok := r.Carts[cartId][productRequest.SKU]; ok {
			item.Count = v.Count
			item.SKU = productRequest.SKU
		}
	}

	return item, nil
}

// AddProductCart - добавляет товар в корзину
func (r *Repository) AddProductCart(productRequest *model.Product, cartId int64) error {
	if _, ok := r.Carts[cartId]; !ok {
		r.Carts[cartId] = make(map[int64]*model.Product)
	}

	r.Carts[cartId][productRequest.SKU] = &model.Product{
		Count: productRequest.Count,
	}

	return nil
}

// DeleteProductCart - удаляет товар из корзины
func (r *Repository) DeleteProductCart(cartId, sku int64) error {
	delete(r.Carts[cartId], sku)
	return nil
}

// ClearCart - чистит корзину
func (r *Repository) ClearCart(cartId int64) error {
	delete(r.Carts, cartId)
	return nil
}

// GetCart - получает содержимое корзины
func (r *Repository) GetCart(cartId int64) (map[int64]*model.Product, error) {
	return r.Carts[cartId], nil
}
