package grpc_handlers

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
)

type Handlers struct {
	loms.UnimplementedLomsServer
	service Service
}

type Service interface {
	Create(context.Context, *model.Order) (model.OrderId, error)
	OrderInfo(context.Context, model.OrderId) (*model.Order, error)
	OrderPay(context.Context, model.OrderId) error
	OrderCancel(context.Context, model.OrderId) error
	StocksInfo(context.Context, uint32) (uint64, error)
}

var (
	requestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "loms_req_total",
			Help: "Loms total amount of request",
		},
		[]string{"handle"},
	)

	requestTimeStatusUrl = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "loms_req_time_status_handle",
			Help:       "Loms summary request time durations second, status, url",
			Objectives: map[float64]float64{.5: .05, .9: .01, .99: .001},
		},
		[]string{"status", "handle"},
	)
	tracer = otel.Tracer(model.ServiceName)
)

func New(service Service) *Handlers {
	return &Handlers{
		service: service,
	}
}
