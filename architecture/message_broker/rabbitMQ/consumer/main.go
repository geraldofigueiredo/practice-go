package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	ampqServerURL := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, err := amqp.Dial(ampqServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	messages, err := channelRabbitMQ.Consume(
		"QueueService1",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("> Received message: %s\n", message.Body)
		}
	}()

	<-forever
}
