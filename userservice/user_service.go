package main

import (
	"Microservices/configuration"
	"Microservices/msgqueue"
	"Microservices/msgqueue/kafka"
	"Microservices/userservice/handler"
	localkafka "Microservices/userservice/kafka"
	"github.com/Shopify/sarama"
)

//func main() {
//	log.Println("Starting server...")
//	router := http.NewServeMux()
//	router.HandleFunc("/",handler.Home)
//	router.HandleFunc("/login",handler.Login)
//	log.Println("Listening on port 8080...")
//	if err := http.ListenAndServe(":8080", router); err != nil {
//		log.Fatal(err)
//	}
//}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var eventListener msgqueue.EventListener
	var eventEmitter msgqueue.EventEmitter

	config, _ :=configuration.ExtractConfiguration("config.json")


	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
	conn, err = sarama.NewClient(config.KafkaMessageBrokers, conf)
	panicIfErr(err)

	eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
	panicIfErr(err)

	eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
	panicIfErr(err)

	processor := localkafka.EventProcessor{EventListener: eventListener}

	go processor.ProcessEvents()

	handler.ServeAPI(config.RestfulEndpoint, eventEmitter)
}