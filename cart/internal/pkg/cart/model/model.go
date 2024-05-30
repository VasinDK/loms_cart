package model

// Корзина
type Cart struct {
	Items      []*Product
	TotalPrice uint32
}

// Товар
type Product struct {
	SKU   int64
	Name  string
	Price uint32
	Count uint16
}
