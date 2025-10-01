package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create a producer: %s\n", err)
		os.Exit(1)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced Events to Topic %s : value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Value))
				}
			}
		}
	}()

	topic := "messages"

	data := "Hello Kafka"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            nil,
		Value:          []byte(data),
	}, nil)

	p.Flush(15 * 1000)
	p.Close()
}
