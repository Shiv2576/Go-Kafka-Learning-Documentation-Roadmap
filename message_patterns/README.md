# Message Patterns

#### Partitions Routing:
Kafka uses the message key to determine which partition the message goes to. By default , Kafka uses a hash of the key to assign the message to a specific partition.

#### Ordering Guarantee:
Kafka guarantees message ordering within the same partition. if all messages with the same key go to the same partition, then those messages will be consumed in a order they are produced.


## Structured Data

Kafka messages have
Key (optional, often a string or ID),Value (the actual payload)

You can serialize structured data in different formats:

```
JSON : Human-readable, easy to use

Avro : Compact, schema evolution, integrates with Schema Registry

Protobuf : Efficient, strong typing, good for cross-language
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
