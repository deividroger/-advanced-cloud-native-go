package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func main()  {

	fmt.Println("Starting RabbitMQ consumer...")
	time.Sleep(10 * time.Second)

	conn, err := amqp.Dial(brokerAddr())
	failOnError(err,"Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err,"Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue(),
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err,"Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"Failed to register a consumer")


	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message, %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}



func  brokerAddr() string {
	brokerAddr := os.Getenv("BROKER_ADDR")

	if len(brokerAddr) == 0 {
		brokerAddr = "amqp://guest:guest@localhost:5672/"

	}
	return brokerAddr

}

func queue() string {
	queue := os.Getenv("QUEUE")

	if len(queue) == 0 {
		queue = "default-queue"
	}
	return  queue
}

func failOnError(err error, msg string)  {

	if err != nil {
		panic(msg)
	}

}