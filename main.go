package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ app started!")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("Connect to RabbitMQ failed: %s ", err)
	}
	defer conn.Close()

	fmt.Println("Successfully Connected To our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Connect to chanel RabbitMQ failed: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("QueueDeclare failed: %s", err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello RabbitMQ now"),
		},
	)

	if err != nil {
		fmt.Printf("ch.Publish failed: %s", err)
	}
	fmt.Println("Successfully Published Message to Queue")
}
