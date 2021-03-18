package builder

import (
	"github.com/nekruzvatanshoev/Microservices/msgqueue"
	"github.com/nekruzvatanshoev/Microservices/msgqueue/kafka"
	"log"
	"os"
)

func NewEventListenerFromEnvironment() (msgqueue.EventListener, error) {
	var listener msgqueue.EventListener
	var err error

	if brokers := os.Getenv("KAFKA_BROKERS"); brokers != "" {
		log.Printf("connecting to Kafka brokers at %s", brokers)

		listener, err = kafka.NewKafkaEventListenerFromEnvironment()
		if err != nil {
			return nil, err
		}
	}

	return listener, nil
}