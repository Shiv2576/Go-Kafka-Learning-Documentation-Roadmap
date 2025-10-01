# Go-Kafka-Learning-Documentation-Roadmap

First install Docker, then pull image for confluent-kafka :

  docker pull confluentinc/cp-kafka:latest

  OR

  brew install confluentinc/tap/cli



to run kafka locally use :

  confluent local kafka start



Or if you want start kafka on a specific local port:

  confluent local kafka start --plaintext-ports 9093

to stop :

  confluent local kafka stop


Getting Started:

  Hello_kafka:
