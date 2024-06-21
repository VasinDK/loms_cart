package model

import "errors"

var ErrOrderIncorrect = errors.New("order Incorrect")                         // не корректный ордер
var ErrSkuNoSuch = errors.New("there is no such sku")                         // sku не существует
var ErrSkuNotEnough = errors.New("SKU is not enough")                         // sku не достаточно
var ErrOrderNoSuch = errors.New("order no such")                              // ордер не существует
var ErrStatusNoAwaitingPayment = errors.New("status no awaiting payment")     // статус не "ожидает оплату"
var ErrStrConnIsEmpty = errors.New("the database connection string is empty") // строка подключения к БД пуста
var ErrDuplicateSku = errors.New("there is a duplicate sku")                  // Имеется дублирование sku
var ErrAddOrder = errors.New("error adding an order")                         // ошибка добавления ордера
var ErrAddStatus = errors.New("error adding status")                          // ошибка установки статуса
