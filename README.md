# Go-Kafka-Learning-Documentation-Roadmap

<p align="center" style="
  background-color: white;
  border-radius: 12px;
  padding: 20px 40px;
  display: inline-block;
  box-shadow: 0 0 10px rgba(0,0,0,0.08);
">
  <img src="assets/golang.svg" alt="Golang" style="width:150px; margin-right:20px;">
  <img src="assets/live.svg" alt="Event Stream" style="width:50px; margin-right:20px;">
  <img src="assets/kafka.svg" alt="Kafka" style="width:150px;">
</p>

<h1 align="center">Go-Kafka Learning Documentation Roadmap</h1>



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
