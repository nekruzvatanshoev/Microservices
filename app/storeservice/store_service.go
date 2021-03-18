package main

import (
	"github.com/nekruzvatanshoev/Microservices/app/storeservice/kafka"

)

func main() {
	kafkaSubscriberConfig := kafka.ConfigureKafkaSubscriber()
	kafka.CreateKafkaSubscriber(kafkaSubscriberConfig)
}