package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/NorseLZJ/example/get_code/config"
)

var (
	cfg = flag.String("conf", "./goGet.json", "go get file")
)

const (
	defProxy = "https:goproxy.cn"
)

func main() {
	cfgT, err := config.Marshal(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatal("GOPATH is nil")
	}

	proxy := os.Getenv("GOPROXY")
	if proxy == "" {
		os.Setenv("GORROXY", defProxy)
	}
	if cfgT.Proxy != "" {
		os.Setenv("GORROXY", cfgT.Proxy)
	}

	codeTotal := len(cfgT.Code)

	wg := sync.WaitGroup{}
	wg.Add(codeTotal)

	for _, v := range cfgT.Code {
		go func(code string) {
			cmd := exec.Command("go", "get", "-u", code)
			err := cmd.Start()
			if err != nil {
				log.Printf("get (%s), Err : %v\n", code, err)
			}
			log.Printf("get (%s) waiting ...\n", code)
			err = cmd.Wait()
			log.Printf("wait err (%s) : %v\n", code, err)
			wg.Done()
		}(v)
	}

	wg.Wait()
}
