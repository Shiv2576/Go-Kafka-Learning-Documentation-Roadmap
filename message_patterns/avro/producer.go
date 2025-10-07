package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/serializers"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	srClient , err := sr.NewClient(sr.Config{
		BaseUrl : "localhost:9093"
	})

	if err != nil {
		fmt.Printf("Failed to create Schema Registry client: %v\n", err)
		os.Exit(1)
	}


	p ,err := kafka.NewProducer(&kafka.ConfigMap{
		"Bootstrap.servers" : "localhost:9093",
		"acks" : "all",
	})

	if err != nil {
		fmt.Printf("failed to create a producer : %s" , err)
		os.Exit(1)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
				case *kafka.Message:
				if ev.TopicPartition.Error !=  nil {
					fmt.Printf("Delivery  failed : %v\n" , ev.TopicPartition.Error )
				} else {
					fmt.Printf("Delivered to topic %s [%d] at offset %v\n", *ev.TopicPartition.Topic , ev.TopicPartition.Partition , ev.TopicPartition.Offset)
				}
			}
		}
	}


	topic := "purchases"
	valuesubject := topic + "-value"


	schemaStr := `{
		"type": "record",
		"name": "Purchase",
		"namespace": "io.confluent.examples",
		"fields": [
			{"name": "user", "type": "string"},
			{"name": "item", "type": "string"},
			{"name": "quantity", "type": "int"},
			{"name": "timestamp", "type": "string"}
		]
	}`

	schemaId , err := srClient.Register(
		valuesubject , sr.Schema{
			Schema:     schemaStr,
			SchemaType: sr.Avro,
		}
	)

	if err != nil {
		fmt.Printf("Failed to register schema: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Using schema ID: %d\n", schemaID)


	avroSerializer := serializers.NewAvroSerializer(srClient , nil)

	users := []string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := []string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10; n++ {
		user := users[rand.Intn(len(users))]
		item := items[rand.Intn(len(items))]
		qty := rand.Intn(5) + 1

		avroRecord := map[string]interface{}{
			"user":      user,
			"item":      item,
			"quantity":  qty,
			"timestamp": time.Now().Format(time.RFC3339),
		}

		valueBytes, err := avroSerializer.Serialize(topic, avroRecord)
		if err != nil {
			fmt.Printf("Serialization error: %v\n", err)
			continue
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(user),
			Value:          valueBytes,
		}, nil)
	}

	p.Flush(1000 * 15)
	p.Close()
}
