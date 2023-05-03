package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Command struct {
	From        int    `json:"from,omitempty"`
	To          int    `json:"to,omitempty"`
	Orbit       int    `json:"orbit,omitempty"`
	Command     string `json:"command"`
	Prompt      bool   `json:"prompt"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Config struct {
	Commands     []Command     `json:"commands"`
	Descriptions []Description `json:"descriptions"`
}

type Description struct {
	Phase       int    `json:"phase"`
	Description string `json:"description"`
}

func ReadConfig(file string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return config, errors.New("unable to read config file")
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, errors.New("unable to parse config file")
	}

	return config, nil
}
