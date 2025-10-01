# Hello Kafka Tutorial (Go + Confluent Kafka)

This is a simple example to understand how to produce and consume messages in Kafka using Go and the [Confluent Kafka Go client](https://github.com/confluentinc/confluent-kafka-go).

---

##  Install Dependencies

```bash
go get github.com/confluentinc/confluent-kafka-go
   Start Kafka with Docker
bash
confluent local kafka start

confluent local kafka topic create messages


Make sure to use the same topic name (messages) in both producer.go and consumer.go.

// Run Producer
bash
Copy code
go build -o out/produce produce.go
./out/produce
// Run Consumer
bash
Copy code
go build -o out/consumer consumer.go
./out/consumer
// Notes on kafka.Message
txt
Copy code
- kafka.Message does not always require key-value pairs.

- You can send only a Value (most common).
- You can send a Key + Value (messages with same key go to the same partition).
- You can send a Key + nil Value (tombstones in compacted topics).
- You can send Key + Value + Headers (for structured metadata).
Examples
go
Copy code
// Only Value
Value: []byte("Hello Kafka")

// Key + Value
Key:   []byte("user123"),
Value: []byte("Hello Kafka")

// Tombstone
Key:   []byte("user123"),
Value: nil

// Key + Value + Headers
Key:   []byte("order123"),
Value: []byte("Order created"),
Headers: []kafka.Header{
    {Key: "correlationId", Value: []byte("abc-123")},
},
📌 Summary
txt
Copy code
- kafka.Message key is optional.
- Key is used for partitioning.
- Value carries the payload.
- Tombstones (key + nil value) are used in compacted topics.
