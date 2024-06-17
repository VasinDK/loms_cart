package model

import "errors"

var ErrOrderIncorrect = errors.New("order Incorrect")                     // не корректный ордер
var ErrSkuNoSuch = errors.New("there is no such sku")                     // sku не существует
var ErrSkuNotEnough = errors.New("SKU is not enough")                     // sku не достаточно
var ErrOrderNoSuch = errors.New("order no such")                          // ордер не существует
var ErrStatusNoAwaitingPayment = errors.New("status no awaiting payment") // статус не "ожидает оплату"
