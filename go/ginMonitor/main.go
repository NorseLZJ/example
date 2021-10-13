package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Listen      string `json:"listen"`
	StopWsGate  string `json:"stop_ws_gate"`
	StartWsGate string `json:"start_ws_gate"`
}

var (
	confPath = flag.String("c", "./config.json", "ginMonitor config path")

	conf = &Config{}
)

func main() {
	flag.Parse()
	DecodeConfig()
	r := gin.Default()
	r.GET("/stopWsGate", stopWsGate)
	r.GET("/startWsGate", startWsGate)
	r.Run(conf.Listen)
}

func stopWsGate(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "stop success"})
	cmd := exec.Command("sh", "-C", conf.StopWsGate)
	cmd.Run()
}

func startWsGate(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "start success"})
	cmd := exec.Command("sh", "-C", conf.StartWsGate)
	cmd.Run()
}

func DecodeConfig() {
	data, err := ioutil.ReadFile(*confPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
}

