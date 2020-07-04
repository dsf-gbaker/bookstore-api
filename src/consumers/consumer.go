package main

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	queue := os.Getenv("RMQ_QNAME")
	uname := os.Getenv("RMQ_UNAME")
	pwd := os.Getenv("RMQ_PWD")
	domain := os.Getenv("RMQ_URL")

	if queue == "" {
		panic("RMQ Queue Name is empty")
	}

	if uname == "" {
		panic("RMQ Username is empty")
	}

	if pwd == "" {
		panic("RMQ Pwd is empty")
	}

	if domain == "" {
		panic("RMQ Domain is empty")
	}

	rmq := fmt.Sprintf("amqp://%s:%s@%s/", uname, pwd, domain)
	conn, err := amqp.Dial(rmq)
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
