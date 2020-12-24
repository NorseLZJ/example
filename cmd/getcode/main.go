package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	goGet = flag.String("cfg", "./goget.json", "go get code file")
	proxy = flag.String("proxy", "https:goproxy.cn", "your proyxy like v2ray host or use default https://goproxy !")
)

type Config struct {
	Code []string `json:"code"`
}

func main() {
	flag.Parse()
	cfgT := &Config{}
	data, err := ioutil.ReadFile(*goGet)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, cfgT)
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("GOPATH") == "" {
		log.Fatal("GOPATH must be set!")
	}

	if os.Getenv("GOPROXY") == "" {
		os.Setenv("GOPROXY", *proxy)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(cfgT.Code))

	for _, v := range cfgT.Code {
		go func(code string) {
			cmd := exec.Command("go", "get", "-u", code)
			err := cmd.Start()
			if err != nil {
				log.Printf("get (%s), Err : %v\n", code, err)
			}
			err = cmd.Wait()
			if err != nil {
				log.Printf("wait err (%s) : %v\n", code, err)
			}
			wg.Done()
		}(v)
	}

	wg.Wait()
}
