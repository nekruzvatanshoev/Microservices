package main

import (
	"github.com/Shopify/sarama"
	"github.com/nekruzvatanshoev/Microservices/app/userservice/handlers"
	localkafka "github.com/nekruzvatanshoev/Microservices/app/userservice/kafka"
	"github.com/nekruzvatanshoev/Microservices/msgqueue"
	"github.com/nekruzvatanshoev/Microservices/msgqueue/configuration"
	"github.com/nekruzvatanshoev/Microservices/msgqueue/kafka"
)

//func main() {
//	log.Println("Starting server...")
//	router := http.NewServeMux()
//	router.HandleFunc("/",handlers.Home)
//	router.HandleFunc("/login",handlers.Login)
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

	config, _ := configuration.ExtractConfiguration("config.json")


	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
	panicIfErr(err)

	eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
	panicIfErr(err)

	eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
	panicIfErr(err)

	processor := localkafka.EventProcessor{EventListener: eventListener}

	go processor.ProcessEvents()

	handlers.ServeAPI(config.RestfulEndpoint, eventEmitter)
}



//func main(){
//	if err := run(); err != nil {
//		log.Println(err)
//		os.Exit(1)
//	}
//}
//
//
//func run() error {
//	return nil
//}