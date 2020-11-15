package main

import (
	"Microservices/storeservice/kafka"

)

func main() {
	kafkaSubscriberConfig := kafka.ConfigureKafkaSubscriber()
	kafka.CreateKafkaSubscriber(kafkaSubscriberConfig)
}