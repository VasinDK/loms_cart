package model

type OrderId int64
type OrderStatus string // статус ордеров

type OrderItem struct {
	Sku   uint32
	Count uint16
}

type Order struct {
	User   int64
	Status OrderStatus
	Items  []*OrderItem
}

type StockItem struct {
	Sku        uint32
	TotalCount uint64
	Reserved   uint64
}

const (
	StatusNew             OrderStatus = "new"
	StatusAwaitingPayment OrderStatus = "awaiting payment"
	StatusFailed          OrderStatus = "failed"
	StatusPayed           OrderStatus = "payed"
	StatusCancelled       OrderStatus = "cancelled"
)

const ServiceName = "Loms"
