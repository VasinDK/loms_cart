package model

import "errors"

var ErrOrderIncorrect = errors.New("order Incorrect")
var ErrSkuNoSuch = errors.New("there is no such sku")
var ErrSkuNotEnough = errors.New("SKU is not enough")
var ErrOrderNoSuch = errors.New("order no such")
var ErrStatusNoAwaitingPayment = errors.New("status no awaiting payment")
