package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type User struct {
	ID     int    `json:"id"`
	User   string `json:"user"`
	Status string `json:"status"`
	Score  int    `json:"score"`
}

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}

			}
		}
	}()

	file, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("Error reading file : ", err)
		return
	}

	var users []User

	if err := json.Unmarshal(file, &users); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	topic := "shopping"

	for _, u := range users {
		data, err := json.Marshal(u)

		if err != nil {
			fmt.Printf("Error encoding users: ", err)
			continue
		}

		key := fmt.Sprintf("%d", u.ID)

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          data,
		}, nil)
	}

	p.Flush(15 * 1000)
	p.Close()
}
