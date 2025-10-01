# Hello Kafka Tutorial (Go + Confluent Kafka)

This is a simple example to understand how to produce and consume messages in Kafka using Go and the [Confluent Kafka Go client](https://github.com/confluentinc/confluent-kafka-go).

---

##  Install Dependencies

First, install the Confluent Kafka Go package:

```bash
go get github.com/confluentinc/confluent-kafka-go
//  Start Kafka with Docker
Make sure you have Docker installed, then start a Kafka broker locally.
For example, using Confluent CLI:

bash
Copy code
confluent local kafka topic create messages
// Ensure the topic name (messages) matches in both producer.go and consumer.go.

 Run the Producer
Build and run the producer that adds messages to the Kafka topic queue:

bash
Copy code
go build -o out/produce produce.go
./out/produce
// Run the Consumer
Build and run the consumer that reads messages from the Kafka topic queue:

bash
Copy code
go build -o out/consumer consumer.go
./out/consumer
// Notes on kafka.Message
In producer.go, the kafka.Message struct does not always require a key.

You can choose what to include depending on your use case:

// Options:
Only Value (most common):

go
Copy code
Value: []byte("Hello Kafka")
Key + Value (used for partitioning):

go
Copy code
Key:   []byte("user123"),
Value: []byte("Hello Kafka")
Tombstone (key with nil value, used in compacted topics):

go
Copy code
Key:   []byte("user123"),
Value: nil
Key + Value + Headers (for structured metadata):

go
Copy code
Key:   []byte("order123"),
Value: []byte("Order created"),
Headers: []kafka.Header{
    {Key: "correlationId", Value: []byte("abc-123")},
},
