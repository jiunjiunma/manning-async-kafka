package service

import "github.com/jiunjiunma/manning-async-kafka/milestone4/model"

type OrderService struct {
	// dummy in-memory map to track previous seen orders
	// in real life, this will be replaced with a db based implementation
	orders map[string]bool
}

func NewOrderService() OrderService {
	return OrderService{
		orders: make(map[string]bool),
	}
}

func (os OrderService) CheckOrderExisted(order model.Order) bool {
	_, ok := os.orders[order.Id]
	return ok
}

func (os OrderService) TrackOrder(order model.Order) error {
	os.orders[order.Id] = true
	return nil
}
