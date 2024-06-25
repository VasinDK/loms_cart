package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/pkg/api/loms/v1"
)

type Config interface {
	GetTokenStore() string
	GetAddressStore() string
}

// Carts - структура: "Корзины". id пользователя - id конкретной корзины.
// Id корзины соответствует первому ключу. SKU товара соответствует
// второму ключу. Значение - продукт
type Carts map[int64]map[int64]*model.Product

// Repository - структура репозитория
type Repository struct {
	Carts      Carts
	ClientLoms loms.LomsClient
	Config     Config
}

// NewRepository - инициализирует репозиторий
func NewRepository(config Config, loms loms.LomsClient) *Repository {
	return &Repository{
		Carts:      make(Carts),
		ClientLoms: loms,
		Config:     config,
	}
}

// CheckSKU - проверяет наличие sku на удаленном сервере
func (r *Repository) CheckSKU(ctx context.Context, ch1 chan<- *model.Product, sku int64) error {
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

	response := &model.Product{}

	jsonBodyCheckSKU, err := json.Marshal(bodyCheckSKU)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		r.Config.GetAddressStore(),
		bytes.NewReader(jsonBodyCheckSKU),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	SkuResponseCheck := CheckSkuResponse{}
	json.NewDecoder(resp.Body).Decode(&SkuResponseCheck)

	// Если в ответе есть цена SKU считаем что товар доступен
	if SkuResponseCheck.Price > 0 {
		response.Name = SkuResponseCheck.Name
		response.Price = SkuResponseCheck.Price
		response.SKU = sku
		ch1 <- response

		return nil
	}

	return model.ErrNoProductInStock
}

// GetProductCart - получает конкретный товар из корзины пользователя
func (r *Repository) GetProductCart(ctx context.Context, productRequest *model.Product, cartId int64) (*model.Product, error) {
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
func (r *Repository) AddProductCart(ctx context.Context, productRequest *model.Product, cartId int64) error {
	if _, ok := r.Carts[cartId]; !ok {
		r.Carts[cartId] = make(map[int64]*model.Product)
	}

	r.Carts[cartId][productRequest.SKU] = &model.Product{
		Count: productRequest.Count,
	}

	return nil
}

// DeleteProductCart - удаляет товар из корзины
func (r *Repository) DeleteProductCart(ctx context.Context, cartId, sku int64) error {
	delete(r.Carts[cartId], sku)
	return nil
}

// ClearCart - чистит корзину
func (r *Repository) ClearCart(ctx context.Context, cartId int64) error {
	delete(r.Carts, cartId)
	return nil
}

// GetCart - получает содержимое корзины
func (r *Repository) GetCart(ctx context.Context, cartId int64) (map[int64]*model.Product, error) {
	return r.Carts[cartId], nil
}

// Checkout - создаент ордера на уделенном сервере
func (r *Repository) Checkout(ctx context.Context, userId int64, cart []*model.Product) (int64, error) {
	items := []*loms.ItemRequest{}

	for i, _ := range cart {
		items = append(items, &loms.ItemRequest{
			Sku:   uint32(cart[i].SKU),
			Count: uint32(cart[i].Count),
		})
	}

	in := &loms.OrderCreateRequest{
		User:  userId,
		Items: items,
	}

	orderIdloms, err := r.ClientLoms.OrderCreate(ctx, in)
	if err != nil {
		return 0, err
	}

	return int64(orderIdloms.GetOrderId()), nil
}

// StockInfo - инфа по остаткам
func (r *Repository) StockInfo(ctx context.Context, sku int64) (int64, error) {
	countloms, err := r.ClientLoms.StocksInfo(ctx, &loms.StocksInfoRequest{Sku: uint32(sku)})

	if err != nil {
		return 0, err
	}

	return int64(countloms.GetCount()), nil
}
