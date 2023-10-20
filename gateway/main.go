package main

import (
	"fmt"
	"grpc-router/config"
	"grpc-router/startup"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Incorrect config")
		return
	}

	gateway := startup.NewServer(conf)
	gateway.Start()
}
