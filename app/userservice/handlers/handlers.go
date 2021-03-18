package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/nekruzvatanshoev/Microservices/app/userservice/domain"
	"github.com/nekruzvatanshoev/Microservices/msgqueue"
	"github.com/nekruzvatanshoev/Microservices/msgqueue/contracts"
	"net/http"
)
//
//func Home(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	fmt.Fprintln(w,"Welcome to Naboft!")
//}
//
//
//func Login(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		w.Header().Set("Content-Type", "application")
//		//var db domain.User
//		db := []domain.User{
//			{
//				ID: "AAABBBCCC",
//				Email: "first.last@company.com",
//				Password: "thelonestarstate",
//				Admin: true,
//			},
//			{
//				ID: "BBBCCCDDD",
//				Email: "first.last@outlook.com",
//				Password: "rockymountainhigh",
//				Admin: false,
//			},
//		}
//
//		fmt.Println(db)
//		var user domain.User
//
//		err := json.NewDecoder(r.Body).Decode(&user)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		marshalUser, err := json.Marshal(user)
//		if err != nil {
//			log.Println(err)
//		}
//		//encodedUser, err := sarama.ByteEncoder.Encode(marshalUser)
//		kafka.CreateKafkaPublisher(marshalUser)
//		fmt.Println(user)
//	}
//
//	return
//
//
//
//}



type CreateLoginHandler struct {
	eventEmitter msgqueue.EventEmitter
}


func (h *CreateLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	//var db domain.User
	db := []domain.User{
		{
			ID: "AAABBBCCC",
			Email: "first.last@company.com",
			Password: "thelonestarstate",
			Admin: true,
		},
		{
			ID: "BBBCCCDDD",
			Email: "first.last@outlook.com",
			Password: "rockymountainhigh",
			Admin: false,
		},
	}

	fmt.Println(db)
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	//encodedUser, err := sarama.ByteEncoder.Encode(marshalUser)


	msg := contracts.UserLoggedInEvent{ID: user.ID}
	h.eventEmitter.Emit(&msg)
}