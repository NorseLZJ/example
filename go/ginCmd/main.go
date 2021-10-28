package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os/exec"
)

type Config struct {
	Listen  string `json:"listen"`
	Options map[string]*struct {
		Start string `json:"start"` // start bat path
		Stop  string `json:"stop"`  // stop bat path
	} `json:"options"`
}

var (
	confPath = flag.String("c", "./config.json", "ginMonitor config path")

	conf = &Config{}
)

func main() {
	flag.Parse()
	DecodeConfig()
	r := gin.Default()
	r.GET("/cmd", execCmd)
	r.Run(conf.Listen)
}

func execCmd(c *gin.Context) {

	group := c.Query("group")
	opt := c.Query("opt")
	opts := conf.Options[group]
	if group == "" || opt == "" || opts == nil {
		c.JSON(200, gin.H{"msg": "cmd error param not enough!"})
		return
	}

	bat := ""
	switch opt {
	case "start":
		bat = opts.Start
	case "stop":
		bat = opts.Stop
	default:
		c.JSON(200, gin.H{"msg": "cmd error param check!"})
		return
	}

	cmd := exec.Command("cmd", "/C", bat)
	cmd.Run()
	c.JSON(200, gin.H{"msg": "cmd success"})
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
	fmt.Println("-----------------------------------------------")
	fmt.Printf("listen:%s\n", conf.Listen)
	for k, v := range conf.Options {
		fmt.Printf("group:%s\nstart:%s\nstop:%s\n", k, v.Start, v.Stop)
	}
	fmt.Println("-----------------------------------------------")
}
