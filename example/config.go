package main

import (
	"encoding/json"

	"github.com/FTIVLTD/vodka"
	"github.com/FTIVLTD/vodka/adapters"

	lib "github.com/niklucky/go-lib"
)

/*
Config - service configuration
*/
type Config struct {
	Version    string           `json:"version"`
	HTTPServer vodka.HTTPConfig `json:"http_server"`
	Postgres   adapters.Config  `json:"postgres"`
	MySQL      adapters.Config  `json:"mysql"`
}

/*
NewConfig - config constructors
*/
func NewConfig(configFileName string) (Config, error) {
	config := Config{}

	fileData, err := lib.ReadFile(configFileName)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
