# Go-Kafka-Learning-Documentation-Roadmap

<p align="center" style="
  background-color: white;
  border-radius: 12px;
  padding: 30px 60px;
  display: inline-flex;
  align-items: center;
  gap: 40px;
  box-shadow: 0 0 10px rgba(0,0,0,0.08);
">
  <img src="assets/golang.svg" alt="Golang" style="width:150px;">
  <img src="assets/live.svg" alt="Event Stream" style="width:50px;">
  <img src="assets/kafka.svg" alt="Kafka" style="width:150px;">
</p>



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
