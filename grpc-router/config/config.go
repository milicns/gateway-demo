package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services map[string]string `yaml:"services"`
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
