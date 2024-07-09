package model

import "errors"

var ErrNoProductInStock = errors.New("no product in stock")
var ErrInsufficientStock = errors.New("insufficient stock")
var ErrOk = errors.New("OK")
var ErrGetPathValueInt = errors.New("error when calling getPathValueInt()")
var ErrValidateVar = errors.New("error when calling validate.Var()")
var ErrValidateStruct = errors.New("error when calling validate.Struct()")
var ErrJsonNewDecoder = errors.New("error when calling json.NewDecoder()")
var ErrJsonMarshal = errors.New("error when calling json.Marshal()")
var ErrHAddProduct = errors.New("error when calling h.AddProduct()")
var ErrHCheckout = errors.New("error when calling h.Checkout()")
var ErrHClearCart = errors.New("error when calling h.ClearCart()")
var ErrHGetCart = errors.New("error when calling h.GetCart()")
var ErrHDeleteProductCart = errors.New("error when calling h.DeleteProductCart()")
