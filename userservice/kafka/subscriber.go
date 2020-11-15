package kafka

import (
	"Microservices/contracts"
	"Microservices/msgqueue"
	"fmt"
	"log"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
}

func (p *EventProcessor) ProcessEvents() {
	log.Println("listening or events")

	received, errors, err := p.EventListener.Listen("userLoggedIn")

	if err != nil {
		panic(err)
	}

	for {
		select {
		case evt := <- received:
			fmt.Printf("got event %T: %s\n", evt, evt)
			p.handleEvent(evt)
		case err = <-errors:
			fmt.Printf("got error while receiving event: %s\n", err)
		}
	}
}


func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.UserLoggedInEvent:
		log.Printf("user %s logged in: %s", e.ID, e)
	default:
		log.Printf("unknown event type: %T", e)
	}
}