package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/nekruzvatanshoev/Microservices/msgqueue"
	"github.com/nekruzvatanshoev/Microservices/msgqueue/contracts"
	"github.com/nekruzvatanshoev/Microservices/business/data/user"
	"net/http"
)


type CreateLoginHandler struct {
	eventEmitter msgqueue.EventEmitter
}


func (h *CreateLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")

	var user user.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	msg := contracts.UserLoggedInEvent{ID: user.ID}
	h.eventEmitter.Emit(&msg)
}