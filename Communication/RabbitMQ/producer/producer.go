package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting with RabbitMQ...")
	time.Sleep(10 * time.Second)

	conn, err := amqp.Dial(brokerAddr())

	failOnError(err,"Failed to conenct to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err,"Failed to conenct to open a channel")
	defer  ch.Close()

	q, err := ch.QueueDeclare(
		queue(),
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err,"Failed to declare a queue")

	msgCount := 0

	doneCh := make(chan struct{})

	go func() {
		for  {
			msgCount++
			body := fmt.Sprintf("Hello RabbitMQ message %v",msgCount)

			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body: []byte(body),
				},
			)


			log.Printf(" [x] sent %s",body)
			time.Sleep(5 * time.Second)


		}
	}()
	<-doneCh
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