package model

import "errors"

var ErrNoProductInStock = errors.New("no product in stock")
var ErrInsufficientStock = errors.New("insufficient stock")
var ErrOk = errors.New("OK")
