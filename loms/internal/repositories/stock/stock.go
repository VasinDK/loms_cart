package stock

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type StockItem struct {
	Sku        uint32 `json:"sku"`
	TotalCount uint64 `json:"total_count"`
	Reserved   uint64 `json:"reserved"`
}

type StockRepository struct {
	Sm ShardManager
}

type ShardManager interface {
	GetShardIndexFromID(id int64) int
	GetShardIndex(key string) (uint32, error)
	Pick(index int) (*pgxpool.Pool, error)
	GetMainShard() int
}

var (
	RequestDBTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "loms_db_total",
			Help: "Loms total amount of request DB",
		},
		[]string{"category"},
	)

	RequestTimeStatusCategoryBD = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "loms_db_time_status_category",
			Help:       "Loms db summary request time durations second, status, category",
			Objectives: map[float64]float64{.5: .05, .9: .05, .99: .05},
		},
		[]string{"status", "category"},
	)
)

// New - создает новый репозиторий для StockRepository
func New(Sm ShardManager) *StockRepository {
	return &StockRepository{
		Sm: Sm,
	}
}
