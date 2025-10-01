# Hello_kafka

Prerequisites Go (version 1.16 or higher) Docker Confluent CLI (optional but recommended for local development) Install Dependencies Install the Confluent Kafka Go client:

```bash
go get github.com/confluentinc/confluent-kafka-go/v2/kafka
```

## Start Kafka Locally with Docker Make sure Docker is installed and running on your system.

```bash
 confluent local kafka start

 confluent local kafka topic create messages
```
Important: The topic name messages must match the topic referenced in both producer.go and consumer.go.



#### Run the Producer The producer sends messages to the messages Kafka topic.

```bash
go build -o out/produce producer.go
./out/produce
```
##### The producer will send a sample message to the topic and exit.


#### Run the Consumer The consumer reads messages from the messages Kafka topic.

```bash
go build -o out/consumer consumer.go
./out/consumer
```
##### The consumer will run continuously and print each received message to the console. Press Ctrl+C to stop it.

Understanding kafka.Message Structure In producer.go, the kafka.Message struct supports several configurations depending on your use case:


```bash
//Value Only (Most Common)

kafka.Message{ Value: []byte("Hello Kafka"), }


//Key and Value (Used for Partitioning)

kafka.Message{ Key: []byte("user123"), Value: []byte("Hello Kafka"), }

//Tombstone Record (Key with Nil Value)

kafka.Message{ Key: []byte("user123"), Value: nil, }

//Key, Value, and Headers (For Metadata)

kafka.Message{ Key: []byte("order123"), Value: []byte("Order created"), Headers: []kafka.Header{ {Key: "correlationId", Value: []byte("abc-123")}, }, }
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
