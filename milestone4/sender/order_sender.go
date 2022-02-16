package sender

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jiunjiunma/manning-async-kafka/milestone4/model"
	"time"
)

var topic = "OrderReceived"

type OrderSender struct {
	p *kafka.Producer
}

func NewSender() (OrderSender, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"acks": "all",
		// don't wait for delivery report
		"go.delivery.reports": false,
	})

	if err != nil {
		return OrderSender{}, err
	}

	return OrderSender{p}, nil
}

func (s OrderSender) SendOrder(order model.Order) error {
	// set timestamp to Now() if not specified
	if order.Timestamp.IsZero() {
		order.Timestamp = time.Now()
	}

	value, err := json.Marshal(order)

	if err != nil {
		return err
	}

	err = s.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: []byte(value)},
		nil,
	)
	return err
}

func (s OrderSender) Close() {
	if s.p != nil {
		s.p.Close()
	}
}
