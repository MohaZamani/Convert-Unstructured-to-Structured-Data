package config

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)


var Config struct {
	// The address to listen on
	Listen string `json:"listen"`
	// The Open AI related keys
	OpenAI struct {
		// The Authorization key
		ApiKey string `json:"api_key"`
		// the temperature hyper paramter for open_ai service
		Temperature float32 `json:"temperature"`
	} `json:"open_ai"`
	// file name for saving structured data to it
	LaptopDetailsFileName string `json:"laptop_details_file_name"`
}

// Load the config file from a location
func LoadConfig(location string) {
	bytes, err := os.ReadFile(location)
	if err != nil {
		log.WithField("error", err).Fatal("Cannot read config file")
	}
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		log.WithField("error", err).Fatal("Cannot parse config file")
	}
}