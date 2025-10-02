# Go-Kafka-Learning-Documentation-Roadmap


## First install Docker, then pull image for confluent-kafka :

```bash
  docker pull confluentinc/cp-kafka:latest

  # or

  brew install confluentinc/tap/cli
```

### To run kafka locally:

```bash
  confluent local kafka start


# if you want start kafka on a specific local port:
  confluent local kafka start --plaintext-ports 9093


# to stop :
  confluent local kafka stop


# to run kafka locally use :
  confluent local kafka

```



## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
