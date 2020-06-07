package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	queue := "TestQueue"

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received: %s\n", d.Body)
		}
	}()

	fmt.Printf("Successfully Connected to RMQ on Queue [%s]\n", queue)
	fmt.Println("[*] - waiting for messages")

	<-forever
}
