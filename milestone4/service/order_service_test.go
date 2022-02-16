package service

import (
	"github.com/jiunjiunma/manning-async-kafka/milestone4/model"
	"testing"
	"time"
)

func TestOrderService(t *testing.T) {
	orderService := NewOrderService()

	order := model.Order {
		"12345",
		"testOrder",
		time.Now(),
		"dummyCustomer",
		[]model.Item{},
	}
	existed := orderService.CheckOrderExisted(order)
	if existed {
		t.Fatalf("Order %v should not in repo", order)
	}

	err := orderService.TrackOrder(order)
	if err != nil {
		t.Fatal("Tracking order failed")
	}
	existed = orderService.CheckOrderExisted(order)
	if !existed {
		t.Fatalf("Order %v should exist in repo after tracking", order)
	}
}
