package model

type Product struct {
	SKU   int64
	Name  string
	Price float32
	Count uint16
}

type CartItem struct {
	Count uint16
}
