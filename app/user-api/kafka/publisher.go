package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func ConfigureKafkaPublisher() *sarama.Config {
	log.Println("Configuring Kafka publisher...")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	return config
}


func CreateKafkaPublisher(data []byte) {
	log.Println("Configuring Kafka publisher...")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	log.Println("Starting Kafka publisher...")

	brokers := []string{"localhost:9092"}
	publisher, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := publisher.Close(); err != nil {
			panic(err)
		}
	}()

	topic := "user-logged-in"
	doneCh := make(chan struct{})
	go func() {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(data),
		}

		partition, offset, err := publisher.SendMessage(msg)
		if err != nil {
			panic(err)
		}

		fmt.Printf("MEssage is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	}()
	<-doneCh
}
