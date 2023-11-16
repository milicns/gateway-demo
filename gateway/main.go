package main

import (
	"fmt"
	"gateway/config"
	"gateway/startup"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	gateway := startup.NewServer(conf)
	gateway.Start()
}
