package main

import (
	"fmt"
	"grpc-router/client"
	"grpc-router/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		return
	}

	clientRegistry := &client.ClientRegistry{
		Clients: make(map[string]client.Client),
	}
	for _, v := range conf.Services {
		clientRegistry.NewClient(v)
	}
	for _, v := range clientRegistry.Clients {
		fmt.Println(v.Cc.Target())
	}

	router := mux.NewRouter().StrictSlash(true)
	c := clientRegistry.Clients["user_service:8080"]
	router.HandleFunc("/user", c.ListServices).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8000"), router))

}
