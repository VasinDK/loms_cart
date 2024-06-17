package grpc_handlers

import (
	"route256/loms/internal/model"
	"route256/loms/pkg/api/loms/v1"
)

// RepackOrderToModel - репак из loms.OrderCreateRequest в model.Order
func (h *Handlers) RepackOrderToModel(order *loms.OrderCreateRequest) (*model.Order, error) {
	var orderModel model.Order

	orderModel.User = order.User
	for i := range order.Items {
		orderModel.Items = append(orderModel.Items, &model.OrderItem{
			Sku:   order.Items[i].Sku,
			Count: uint16(order.Items[i].Count),
		})
	}

	return &orderModel, nil
}

// RepackOrderToProto - репак из model.Order в loms.OrderInfoResponse
func (h *Handlers) RepackOrderToProto(order *model.Order) (*loms.OrderInfoResponse, error) {
	orderProto := &loms.OrderInfoResponse{
		Status: string(order.Status),
		User:   order.User,
	}

	for i := range order.Items {
		orderProto.Items = append(orderProto.Items, &loms.ItemRequest{
			Sku:   order.Items[i].Sku,
			Count: uint32(order.Items[i].Count),
		})
	}

	return orderProto, nil
}

// RepackOrderIdToProto - репак из model.OrderId в loms.OrderId
func (h *Handlers) RepackOrderIdToProto(orderIdModel model.OrderId) *loms.OrderId {
	orderId := &loms.OrderId{
		OrderId: int64(orderIdModel),
	}

	return orderId
}
