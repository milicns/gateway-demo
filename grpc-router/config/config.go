package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services     map[string]ServiceConfig `yaml:"services"`
	GatewayRoute string                   `yaml:"gateway_route"`
}

type ServiceConfig struct {
	Address      string                  `yaml:"address"`
	ServiceRoute string                  `yaml:"service_route"`
	Methods      map[string]MethodConfig `yaml:"grpc_methods"`
}

type MethodConfig struct {
	MethodRoute string `yaml:"method_route"`
	Type        string `yaml:"type"`
}

func LoadConfig() (*Config, error) {
	path := "config.yml"

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
