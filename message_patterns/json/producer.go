package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// This producer sends structured messages in Json Format:
type PurchaseEvent struct {
	User      string `json:"user"`
	Item      string `json:"item"`
	Quantity  int    `json:"quantitiy"`
	Timestamp string `json:"timestamp"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9093",
			"acks":              "all",
		})

	if err != nil {
		fmt.Printf("Failed to create producer : %s", err)
		os.Exit(1)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery Failed : %v\n", ev.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered to topic %s [%d] at offset %v\n", *ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}
	topic := "purchases"

	for n := 0; n < 10; n++ {
		user := users[rand.Intn(len(users))]
		item := items[rand.Intn(len(items))]
		qty := rand.Intn(5) + 1

		event := PurchaseEvent{
			User:      user,
			Item:      item,
			Quantity:  qty,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		jsonValue, err := json.Marshal(event)

		if err != nil {
			fmt.Printf("Failed to marshal JSON: %v\n", err)
			continue
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(user),
			Value:          jsonValue,
		}, nil)
	}

	p.Flush(5 * 1000)
	p.Close()
}
