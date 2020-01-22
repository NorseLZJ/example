package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GetConfig struct {
	Code  []string `json:"code"`
	Proxy string   `json:"proxy"`
}

func getContext(file string) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(fp)
}

func Marshal(file string, ret interface{}) error {
	data, err := getContext(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, ret)
	return err
}
