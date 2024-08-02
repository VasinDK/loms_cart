package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"route256/cart/internal/model"
	"route256/cart/pkg/api/loms/v1"
	"route256/cart/pkg/statuses"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/redis/go-redis/v9"
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
	InMemoryDB *redis.Client
	mu         sync.RWMutex
}

var (
	requestOutTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cart_out_req_total",
			Help: "Total out amount of request ",
		},
		[]string{"externalAddress"},
	)

	requestOutTimeStatusAddress = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "cart_out_request_time_status_category",
			Help:       "Cart out summary request time durations second, status, url",
			Objectives: map[float64]float64{.5: .05, .9: .05, .99: .05},
		},
		[]string{"status", "externalAddress"},
	)
)

// NewRepository - инициализирует репозиторий
func NewRepository(config Config, loms loms.LomsClient, inMemoryDB *redis.Client) *Repository {
	return &Repository{
		InMemoryDB: inMemoryDB,
		Carts:      make(Carts),
		ClientLoms: loms,
		Config:     config,
	}
}

// CheckSKU - проверяет наличие sku на удаленном сервере
func (r *Repository) CheckSKU(ctx context.Context, sku int64) (*model.Product, error) {
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
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		r.Config.GetAddressStore(),
		bytes.NewReader(jsonBodyCheckSKU),
	)

	if err != nil {
		return nil, err
	}

	start := time.Now()
	externalAddress := r.Config.GetAddressStore()

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	requestOutTotal.WithLabelValues(externalAddress).Inc()
	requestOutTimeStatusAddress.WithLabelValues(statuses.GetCodeHTTP(err), externalAddress).Observe(time.Since(start).Seconds())

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	SkuResponseCheck := CheckSkuResponse{}
	json.NewDecoder(resp.Body).Decode(&SkuResponseCheck)

	// Если в ответе есть цена SKU считаем что товар доступен
	if SkuResponseCheck.Price > 0 {
		response.Name = SkuResponseCheck.Name
		response.Price = SkuResponseCheck.Price
		response.SKU = sku

		return response, nil
	}

	return nil, model.ErrNoProductInStock
}

// GetProductCart - получает конкретный товар из корзины пользователя
func (r *Repository) GetProductCart(ctx context.Context, productRequest *model.Product, cartId int64) (*model.Product, error) {
	item := &model.Product{}

	r.mu.RLock()
	v, ok := r.Carts[cartId][productRequest.SKU]
	r.mu.RUnlock()

	if ok {
		item.Count = v.Count
		item.SKU = productRequest.SKU
	}

	return item, nil
}

// AddProductCart - добавляет товар в корзину
func (r *Repository) AddProductCart(ctx context.Context, productRequest *model.Product, cartId int64) error {

	r.mu.RLock()
	_, ok := r.Carts[cartId]
	r.mu.RUnlock()

	if !ok {
		r.mu.Lock()
		r.Carts[cartId] = make(map[int64]*model.Product)
		r.mu.Unlock()
	}

	r.mu.Lock()
	r.Carts[cartId][productRequest.SKU] = &model.Product{
		Count: productRequest.Count,
	}
	r.mu.Unlock()

	return nil
}

// DeleteProductCart - удаляет товар из корзины
func (r *Repository) DeleteProductCart(ctx context.Context, cartId, sku int64) error {
	r.mu.Lock()
	delete(r.Carts[cartId], sku)
	r.mu.Unlock()

	return nil
}

// ClearCart - чистит корзину
func (r *Repository) ClearCart(ctx context.Context, cartId int64) error {
	r.mu.Lock()
	delete(r.Carts, cartId)
	r.mu.Unlock()

	return nil
}

// GetCart - получает содержимое корзины
func (r *Repository) GetCart(ctx context.Context, cartId int64) (map[int64]*model.Product, error) {
	r.mu.RLock()
	cart := r.Carts[cartId]
	r.mu.RUnlock()

	return cart, nil
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

	start := time.Now()
	externalAddress := "ClientLoms.OrderCreate"

	orderIdloms, err := r.ClientLoms.OrderCreate(ctx, in)

	requestOutTotal.WithLabelValues(externalAddress).Inc()
	requestOutTimeStatusAddress.WithLabelValues(statuses.GetStatusCodeGRPC(err), externalAddress).Observe(time.Since(start).Seconds())

	if err != nil {
		return 0, err
	}

	return int64(orderIdloms.GetOrderId()), nil
}

// StockInfo - инфа по остаткам
func (r *Repository) StockInfo(ctx context.Context, sku int64) (int64, error) {
	start := time.Now()
	externalAddress := "ClientLoms.StocksInfo"

	countloms, err := r.ClientLoms.StocksInfo(ctx, &loms.StocksInfoRequest{Sku: uint32(sku)})

	requestOutTotal.WithLabelValues(externalAddress).Inc()
	requestOutTimeStatusAddress.WithLabelValues(statuses.GetStatusCodeGRPC(err), externalAddress).Observe(time.Since(start).Seconds())

	if err != nil {
		return 0, err
	}

	return int64(countloms.GetCount()), nil
}

// GetByKeyMemDB - получает значение по ключу из inMemoryDB
func (r *Repository) GetByKeyMemDB(ctx context.Context, key string) (string, error) {
	start := time.Now()
	externalAddress := "GetByKeyMemDB"

	val, err := r.InMemoryDB.Get(ctx, key).Result()

	requestOutTotal.WithLabelValues(externalAddress).Inc()
	requestOutTimeStatusAddress.WithLabelValues(statuses.GetStatusCodeRedis(err), externalAddress).Observe(time.Since(start).Seconds())

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

// SetKeyMemDB - устанавливает ключ - значение в inMemoryDB
func (r *Repository) SetKeyMemDB(ctx context.Context, key string, value string) error {
	start := time.Now()
	externalAddress := "SetKeyMemDB"

	err := r.InMemoryDB.Set(ctx, key, value, 0).Err()

	requestOutTotal.WithLabelValues(externalAddress).Inc()
	requestOutTimeStatusAddress.WithLabelValues(statuses.GetStatusCodeRedis(err), externalAddress).Observe(time.Since(start).Seconds())

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ConnWebSocket(w http.ResponseWriter, req *http.Request, BufferSize int) (*websocket.Conn, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  BufferSize,
		WriteBufferSize: BufferSize,
	}

	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		return nil, fmt.Errorf("upgrader.Upgrade %w", err)
	}

	return ws, nil
}

func (r *Repository) ReadWebSocket() {

}

func (r *Repository) WriteWebSocket() {

}
