package handler

import (
	"Microservices/userservice/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w,"Welcome to Naboft!")
}


func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

		fmt.Println(user)
	}

	return



}