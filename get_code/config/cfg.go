package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GetConfig struct {
	Get   []string `json:"get"`
	Proxy string   `json:"proxy"`
}

func getContext(file string) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(fp)
}

func Marshal(file string) (*GetConfig, error) {
	data, err := getContext(file)
	if err != nil {
		return nil, err
	}

	cfg := &GetConfig{}
	err = json.Unmarshal(data, cfg)
	return cfg, err
}
