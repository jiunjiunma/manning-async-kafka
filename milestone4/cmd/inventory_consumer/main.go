package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jiunjiunma/manning-async-kafka/milestone4/model"
	"github.com/jiunjiunma/manning-async-kafka/milestone4/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	broker := "localhost"
	group := "inventory-consumer"
	topics := []string {"OrderReceived"}
	errorTopic := "Error"
	confirmedTopic := "OrderConfirmed"
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		// Avoid connecting to IPv6 brokers:
		// This is needed for the ErrAllBrokersDown show-case below
		// when using localhost brokers on OSX, since the OSX resolver
		// will return the IPv6 addresses first.
		// You typically don't need to specify this configuration property.
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "latest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"acks": "all",
		// don't wait for delivery report
		"go.delivery.reports": false,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create producer: %s\n", err)

		c.Close()
		os.Exit(1)
	}

	orderService := service.NewOrderService()

	err = c.SubscribeTopics(topics, nil)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				var order model.Order
				err = json.Unmarshal(e.Value, &order)
				if err != nil {
					fmt.Printf("Error, unable to convert message %s to order.\n", string(e.Value))
					continue
				}
				if orderService.CheckOrderExisted(order) {
					fmt.Fprintf(os.Stderr, "Error, duplicate order %s, send to error topic.\n", string(e.Value))
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &errorTopic, Partition: kafka.PartitionAny},
						Value: e.Value},
						nil,
					)
				} else {
					// success, track order then send to OrderConfirmed
					orderService.TrackOrder(order)
					fmt.Printf("Success, not a duplicate order, send event to topic ConfirmedOrder.\n")
					p.Produce(&kafka.Message{
						TopicPartition: kafka.TopicPartition{Topic: &confirmedTopic, Partition: kafka.PartitionAny},
						Value:          e.Value},
						nil,
					)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

	fmt.Printf("Closing producer\n")
	p.Close()
}