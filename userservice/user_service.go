package main

import (
	"log"
	"net/http"
	"Microservices/userservice/handler"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/",handler.Home)
	router.HandleFunc("login",handler.Login)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}