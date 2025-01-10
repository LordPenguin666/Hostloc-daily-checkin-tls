package config

import (
	"encoding/json"
	"go.uber.org/zap"
	"os"
)

func ReadConfig(fileName string, log *zap.Logger) *Config {
	config := &Config{}

	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = json.Unmarshal(b, config); err != nil {
		log.Fatal(err.Error())
	}

	return config
}
