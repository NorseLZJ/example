package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	SaveDir string `json:"saveDir"`
	Port    string `json:"port"`
	Key     string `json:"key"`
}

var (
	cfg *Config
)

func readConfig() {
	cfg = &Config{}
	data, err := ioutil.ReadFile(*config)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		panic(err)
	}
}
