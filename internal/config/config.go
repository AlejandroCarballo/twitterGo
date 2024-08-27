package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	Host         string
	Port         string
	DatabaseFile string
	HostKeyPath  string
}

func NewConfig() (*Config, error) {
	var cf *Config
	confFileDir, _ := os.Getwd()
	confFile, err := os.Open(path.Join(confFileDir, "config.json"))
	if err != nil {
		return nil, err
	}

	defer confFile.Close()

	jsonParser := json.NewDecoder(confFile)
	decErr := jsonParser.Decode(&cf)

	if decErr != nil {
		return nil, err
	}

	return cf, nil
}
