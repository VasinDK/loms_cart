package http_handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

var (
	requestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cart_req_total",
			Help: "Total amount of request",
		},
		[]string{"url"},
	)

	requestTimeStatusUrl = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "cart_req_time_status_url",
			Help:       "Cart summary request time durations second, status, url",
			Objectives: map[float64]float64{.5: .05, .9: .01, .99: .001},
		},
		[]string{"status", "url"},
	)
)

// New - инициализирует экземпляр транспорта
func New() *Server {
	return &Server{}
}

var validate = validator.New(validator.WithRequiredStructEnabled())
