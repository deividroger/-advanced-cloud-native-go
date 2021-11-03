package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting synchronous kafka producer...")
	time.Sleep( 5 * time.Second)

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	brokers := []string{brokerAddr()}
	producer, err := sarama.NewSyncProducer(brokers,config)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err:= producer.Close(); err != nil{
			panic(err)
		}
	}()

	topic := topic()
	msgCount := 0

	doneCh := make(chan struct{})

	go func() {
		for  {
			msgCount++

			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(fmt.Sprintf("Hello Kafka %v", msgCount)),
			}

			_, _, err := producer.SendMessage(msg)

			if err != nil {
				panic(err)
			}

			time.Sleep( 5 * time.Second)
		}
	}()

	 <-doneCh
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