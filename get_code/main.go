package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/example/get_code/config"
)

var (
	cfg = flag.String("conf", "./goGet.json", "go get file")
)

const (
	defProxy = "https:goproxy.cn"
)

func main() {

	proxy := os.Getenv("GOPROXY")
	if proxy == "" {
		os.Setenv("GORROXY", defProxy)
	}

	cfgT, err := config.Marshal(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	codeTotal := len(cfgT.Code)

	wg := sync.WaitGroup{}
	wg.Add(codeTotal)

	for _, v := range cfgT.Code {
		go func(code string) {
			cmd := exec.Command("go", "get", "-u", code)
			err := cmd.Start()
			if err != nil {
				log.Printf("get (%s), err : %v\n", code, err)
			}
			log.Printf("wait get (%s) finish\n", code)
			err = cmd.Wait()
			log.Printf("get (%s) err: %v\n", code, err)
			wg.Done()
		}(v)
	}

	wg.Wait()
}
