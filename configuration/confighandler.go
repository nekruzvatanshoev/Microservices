package configuration



var (
	RestfulEPDefault = "localhost:8181"
	KafkaMessageBrokersDefault = []string{"localhost:9092"}
)


type ServiceConfig struct {
	RestfulEndpoint string `json:"restful_endpoint"`
	KafkaMessageBrokers []string `json:"kafka_message_brokers"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		RestfulEndpoint: RestfulEPDefault,
		KafkaMessageBrokers: KafkaMessageBrokersDefault,
	}

	return conf, nil
}