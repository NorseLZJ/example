package main

import (
	"encoding/json"
	"log"
	"os"
)

type RedisConf struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

type Config struct {
	Mysql  []string  `json:"mysql"`
	Redis  RedisConf `json:"redis"`
	JwtKey string    `json:"jwt_key"`
	Listen string    `json:"listen"`
}

var (
	GConfig = Config{}
)

func ConfInit() {
	data, err := os.ReadFile(*config)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &GConfig); err != nil {
		log.Fatal(err)
	}
}
