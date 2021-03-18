package main

import (
	"github.com/nekruzvatanshoev/Microservices/storeservice/kafka"

)

func main() {
	kafkaSubscriberConfig := kafka.ConfigureKafkaSubscriber()
	kafka.CreateKafkaSubscriber(kafkaSubscriberConfig)
}