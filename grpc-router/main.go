package main

import (
	"context"
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
		fmt.Println(err)
		return
	}

	clientRegistry := &client.ClientRegistry{
		Clients: make(map[string]client.Client),
	}

	for _, s := range conf.Services {
		clientRegistry.NewClient(s.Address, s.Methods)
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, s := range conf.Services {
		c := clientRegistry.Clients[s.Address]
		for mtdName, mtdConf := range s.Methods {
			path := conf.GatewayRoute + s.ServiceRoute + mtdConf.MethodRoute
			fmt.Println(path, mtdConf.Type)
			router.HandleFunc("/"+path, middleware(mtdName, c.InvokeGrpcMethod)).Methods(mtdConf.Type)
			fmt.Println(router.GetRoute("/" + path))
		}
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8000"), router))

}
func middleware(mtdName string, h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "mtdName", mtdName)
		h.ServeHTTP(w, r.WithContext(ctx))
	})

}
