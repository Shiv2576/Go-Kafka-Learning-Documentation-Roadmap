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


## Schema Registry

schema registry is not a part of kafka but is used as an add-on that helps manage schemas for structured data
(like avro , Protobuf , Json Format)

Schema registry is seperate service (usually running alongside kafka) that :
1. Stores Avro (or Protobuf/JSON Schema) definitions.
2. Assigns each schema a unique ID.
3. When a producer sends a message:
    It looks up or registers the schema in Schema Registry.
    It gets a schema ID.
    It sends the message as:
        [MAGIC_BYTE (1 byte)] + [SCHEMA_ID (4 bytes)] + [AVRO_BINARY_PAYLOAD]


4. When a consumer receives the message:
    It reads the schema ID from the first 5 bytes.
    It fetches the schema from Schema Registry using that ID.
    It uses the schema to decode the binary Avro payload.


### Why do we need schema registry ?

1. Prevent broken data – Enforces a valid schema so producers can’t send malformed messages.
2. Safe schema evolution – Lets you add/change fields without breaking existing consumers (via compatibility rules).
3. Team contract – Acts as a shared source of truth for data format across teams and services.
4. Compact binary format – Avro + Schema Registry = smaller messages → less storage & network usage.
5. Centralized governance – View, manage, and version all schemas in one place (with UI/API).

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
