package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type User struct {
	ID     int    `json:"id"`
	User   string `json:"user"`
	Status string `json:"status"`
	Score  int    `json:"score"`
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"group.id":          "stream-processing",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer : %s", err)
		os.Exit(1)
	}

	topic := "shopping"
	err = c.SubscribeTopics([]string{topic}, nil)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {

	}

}
