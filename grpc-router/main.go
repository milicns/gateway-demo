package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	userRefClient, userCc := NewUserClient("user_service:8080")
	orderRefClient, orderCc := NewOrderClient("order_service:8080")
	userInvoker := UserInvoker{refClient: userRefClient, cc: userCc}
	orderInvoker := OrderInvoker{refClient: orderRefClient, cc: orderCc}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", userInvoker.invokeCreateUser).Methods("POST")
	router.HandleFunc("/order", orderInvoker.invokeCreateOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("GRPC_ROUTER_PORT")), router))

}
