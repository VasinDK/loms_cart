package model

// Cart - корзина
type Cart struct {
	Items      []*Product
	TotalPrice uint32
}

// Product - товар
type Product struct {
	SKU   int64
	Name  string
	Price uint32
	Count uint16
}

const ServiceName = "Cart"