package stock

import (
	_ "embed"
	"encoding/json"
	"log/slog"
	"os"
)

type StockItem struct {
	Sku        uint32 `json:"sku"`
	TotalCount uint64 `json:"total_count"`
	Reserved   uint64 `json:"reserved"`
}

type StockRepository struct {
	Repo map[uint32]StockItem
}

//go:embed stock-data.json
var byteJSON []byte

// New - создает новый репозиторий для StockRepository
func New() *StockRepository {
	const op = "StockRepository.New"

	var Stocks []StockItem

	err := json.Unmarshal(byteJSON, &Stocks)
	if err != nil {
		slog.Error(op, "json.Unmarshal", err.Error())
		os.Exit(1)
	}

	Stock := &StockRepository{
		Repo: make(map[uint32]StockItem),
	}

	for i := range Stocks {
		Stock.Repo[Stocks[i].Sku] = Stocks[i]
	}

	return Stock
}
