package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	sr "github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"

	"github.com/confluentinc/confluent-kafka-go/v2/serializers"
)

func main() {

	srClient, err := sr.NewClient(sr.Config{
		BaseURL: "http://localhost:8081",
	})

	if err != nil {
		fmt.Printf("Failed to create Schema Registry client: %v\n", err)
		os.Exit(1)
	}
	avroDeserializer := serializers.NewAvroDeserializer(srClient, nil)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"group.id":          "avro",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "purchases"
	err = c.SubscribeTopics([]string{topic}, nil)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("caught signal %v : terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}

			record, desErr := avroDeserializer.Deserialize(topic, ev.Value)

			if desErr != nil {
				fmt.Printf("❌ Avro deserialization failed: %v\n", desErr)
				continue
			}

			fmt.Printf("Consumed event from topic %s: key = %-10s value = %+v\n",
				*ev.TopicPartition.Topic,
				string(ev.Key),
				record)
		}
	}

	c.Close()

}
