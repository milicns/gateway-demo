package config

import "os"

type Config struct {
	Port         string
	OrdersDBHost string
	OrdersDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:         os.Getenv("USER_SERVICE_PORT"),
		OrdersDBHost: os.Getenv("ORDERS_DB_HOST"),
		OrdersDBPort: os.Getenv("ORDERS_DB_PORT"),
	}
}
