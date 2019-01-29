package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	AccessID  string
	AccessKey string
}

func GetConfigs() (Config, error) {
	configuration := Config{}
	_, path, _, _ := runtime.Caller(0)

	path, _ = filepath.Abs(filepath.Dir(path))

	file, err := os.Open(path + "/config.json")

	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)
	return configuration, err

}
