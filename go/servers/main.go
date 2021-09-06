package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

type Config struct {
	Configs []struct {
		Path  string   `json:"path"`
		Param []string `json:"param"`
	} `json:"configs"`
}

var (
	configFile = flag.String("c", "./config.json", "servers config.")

	config = &Config{}
)

func readConfig() {
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}
	config = &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
}

func Task(idx int, stop chan int, wg *sync.WaitGroup) {

	tc := config.Configs[idx]
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, tc.Path, tc.Param...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	// cmd.Stdout = os.Stdout

	wg.Add(1)

	go func() {
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Cmd Run err :%s\n", err.Error())
			return
		}
		for {
			select {
			case i, ok := <-stop:
				if ok && i == 1 {
					cancel()
					err = cmd.Wait()
					wg.Done()
					return
				}
			}
		}
	}()
}

func main() {
	flag.Parse()
	readConfig()
	stop := make(chan int)
	wg := &sync.WaitGroup{}
	for idx := range config.Configs {
		Task(idx, stop, wg)
	}
	time.Sleep(time.Second * 10)
	num := len(config.Configs)
	for i := 0; i < num; i++ {
		stop <- 1
	}
	wg.Wait()
}
