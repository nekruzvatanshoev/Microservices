package msgqueue

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nekruzvatanshoev/Microservices/msgqueue/contracts"
	"log"
)

type StaticEventMapper struct {}


func (e *StaticEventMapper) MapEvent(eventName string, serialized interface{}) (Event, error) {

	var event Event
	log.Println("EVENT NAME:",eventName)
	switch eventName {
	case "userLoggedIn":
		event = &contracts.UserLoggedInEvent{}
	default:
		return nil, fmt.Errorf("unknown event type %s", eventName)
	}

	switch s := serialized.(type) {

	case []byte:
		err := json.Unmarshal(s, event)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s",eventName, err)
		}

	default:
		cfg := mapstructure.DecoderConfig{
			Result: event,
			TagName: "json",
		}

		dec ,err := mapstructure.NewDecoder(&cfg)
		if err != nil {
			return nil, fmt.Errorf("could not initialize decoder for event %s: %s", eventName, err)
		}

		err = dec.Decode(s)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal event %s: %s", eventName, err)
		}
	}


	return event, nil
}