package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type Order struct {
	Id int64 `json:"id"`
	Name string	`json:"name"`
	Timestamp int64	`json:"timestamp"`
	Body string	`json:"body"`
}

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		//"go.delivery.reports": false,
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "OrderReceived"
	orders := []Order {
		Order{1, "order1", time.Now().UnixMilli(), "order1 details"},
		Order{2, "order2", time.Now().UnixMilli(), "order2 details"},
		Order{3, "order2", time.Now().UnixMilli(), "order3 details"},
	}
	for _, order := range orders {
		message, err := json.Marshal(order)
		if err != nil {
			fmt.Println("Marshaling failed for order", order.Id)
			continue
		}
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          message,
		}, nil)
	}

	// Wait for message deliveries before shutting down
	for notDelivered := p.Flush(1000); notDelivered > 0; notDelivered = p.Flush(1000) {
		fmt.Println(notDelivered, "messages not delivered yet")
	}
	fmt.Println("Done! all messages delivered")
}