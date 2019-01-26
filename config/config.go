package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AccessID  string
	AccessKey string
}

func GetConfigs() (Config, error) {
	configuration := Config{}
	file, err := os.Open("config/config.json")
	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)
	return configuration, err

}
