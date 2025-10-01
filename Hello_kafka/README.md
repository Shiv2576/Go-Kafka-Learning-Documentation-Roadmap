#Hello Kafka Tutorial (Go + Confluent Kafka)


Prerequisites
Go (version 1.16 or higher)
Docker
Confluent CLI (optional but recommended for local development)
Install Dependencies
Install the Confluent Kafka Go client:

bash

go get github.com/confluentinc/confluent-kafka-go/v2/kafka

Start Kafka Locally with Docker
Make sure Docker is installed and running on your system.

If you are using the Confluent CLI, start a local Kafka cluster and create a topic:

bash

confluent local kafka start
confluent local kafka topic create messages

Important: The topic name messages must match the topic referenced in both producer.go and consumer.go.

If you are not using the Confluent CLI, you can use any local Kafka setup (e.g., docker-compose with ZooKeeper and Kafka) as long as the broker is accessible at localhost:9092.

Run the Producer
The producer sends messages to the messages Kafka topic.

Build and run the producer:

bash

go build -o out/produce producer.go
./out/produce
The producer will send a sample message to the topic and exit.

Run the Consumer
The consumer reads messages from the messages Kafka topic.

Build and run the consumer:

bash

go build -o out/consumer consumer.go
./out/consumer
The consumer will run continuously and print each received message to the console. Press Ctrl+C to stop it.

Understanding kafka.Message Structure
In producer.go, the kafka.Message struct supports several configurations depending on your use case:

1. Value Only (Most Common)

kafka.Message{
    Value: []byte("Hello Kafka"),
}


2. Key and Value (Used for Partitioning)

kafka.Message{
    Key:   []byte("user123"),
    Value: []byte("Hello Kafka"),
}


3. Tombstone Record (Key with Nil Value)

kafka.Message{
    Key:   []byte("user123"),
    Value: nil,
}


4. Key, Value, and Headers (For Metadata)

kafka.Message{
    Key:   []byte("order123"),
    Value: []byte("Order created"),
    Headers: []kafka.Header{
        {Key: "correlationId", Value: []byte("abc-123")},
    },
}
