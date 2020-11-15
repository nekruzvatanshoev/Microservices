package main

import (
	"log"
	"net/http"
	"Microservices/userservice/handler"
)

func main() {
	log.Println("Starting server...")
	router := http.NewServeMux()
	router.HandleFunc("/",handler.Home)
	router.HandleFunc("/login",handler.Login)
	log.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}