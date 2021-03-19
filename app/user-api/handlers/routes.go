package handlers

import (
	"github.com/gorilla/mux"
	"github.com/nekruzvatanshoev/Microservices/msgqueue"
	"net/http"
	"time"
)

func ServeAPI(listenAddr string, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	r.Methods("post").Path("/login").Handler(&CreateLoginHandler{eventEmitter: eventEmitter})

	srv := http.Server {
		Handler: r,
		Addr: listenAddr,
		WriteTimeout: 2 *time.Second,
		ReadTimeout: 1 * time.Second,
	}

	srv.ListenAndServe()
}
