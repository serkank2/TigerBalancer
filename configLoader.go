package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	ListenerProtocol string `yaml:"ListenerProtocol"`
}

func ConfigLoad() *Config {
	yFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("Error reading config file:  ", err)
		os.Exit(1)
	}
	var config Config
	err = yaml.Unmarshal([]byte(yFile), &config)
	if err != nil {
		log.Fatal("Failed to unmarshall:", err)
		os.Exit(1)
	}
	return &config
}
