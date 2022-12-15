package loader

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"tigerblancer/model"
)

func ConfigLoad() *model.Config {
	yFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("Error reading config file:  ", err)
		os.Exit(1)
	}
	var config model.Config
	err = yaml.Unmarshal([]byte(yFile), &config)
	if err != nil {
		log.Fatal("Failed to unmarshall:", err)
		os.Exit(1)
	}
	return &config
}
