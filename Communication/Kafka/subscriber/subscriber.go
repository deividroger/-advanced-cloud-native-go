package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Starting syncronous kafka subscriber...")
	time.Sleep(5 * time.Second)


	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.CommitInterval = 5 * time.Second
	config.Consumer.Return.Errors = true


	brokers := []string{brokerAddr()}

	master, err := sarama.NewConsumer(brokers,config)

	if err != nil{
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()


	consumer,err := master.ConsumePartition(topic(),0,sarama.OffsetOldest)

	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal,1)
	signal.Notify(signals,os.Interrupt)

	msgCount := 0
	doneCh := make(chan struct{})

	go func() {
		for  {
			select {
				case err := <- consumer.Errors():
					fmt.Println(err)
				case msg := <- consumer.Messages():
					msgCount++
					fmt.Println("Received messages", string(msg.Key),string(msg.Value))
				case <- signals:
					fmt.Println("Interrupt is detected")
					doneCh <- struct{}{}
			}
		}
	}()
	<- doneCh
	fmt.Println(msgCount, "messages")
}
func topic() string {
	topic := os.Getenv("TOPIC")

	if len(topic) == 0 {
		topic = "default-topic"
	}
	return  topic
}

func brokerAddr() string {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if len(brokerAddr) == 0 {
		brokerAddr = "localhost:9082"
	}
	return  brokerAddr
}