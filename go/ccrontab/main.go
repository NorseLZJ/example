package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
)

type Task struct {
	Cmd    string `json:"cmd"`    // sh
	Param  string `json:"param"`  // -c
	Script string `json:"script"` // /aaa/bbb.sh
	Space  string `json:"space"`  // * * * * * *
}

type Config struct {
	Tasks []Task `json:"tasks"`
}

var (
	cfg    = flag.String("c", "./ccb.json", "crontab config json")
	logDir = flag.String("l", "./ccb.log", "crontab log")
)

func main() {
	flag.Parse()
	cfgT := &Config{}
	data, err := ioutil.ReadFile(*cfg)
	exitErr(err)
	err = json.Unmarshal(data, cfgT)
	exitErr(err)
	fd, err := os.OpenFile(*logDir, os.O_WRONLY|os.O_TRUNC, 0600)
	exitErr(err)
	defer fd.Close()
	log.SetOutput(fd)
	c := cron.New(cron.WithSeconds())
	for _, vv := range cfgT.Tasks {
		_, err = c.AddFunc(vv.Space, func() {
			script, err := ioutil.ReadFile(vv.Script)
			exitErr(err)
			cmd := exec.Command(vv.Cmd, vv.Param, string(script))
			b, err := cmd.Output()
			exitErr(err)
			fd.Write([]byte(fmt.Sprintf("\nCCB START TIME(%s)\n", time.Now().String())))
			fd.Write(b)
			fd.Write([]byte(fmt.Sprintf("\nCCB END TIME(%s)\n", time.Now().String())))
		})
		exitErr(err)
	}
	c.Start()
	waitExit()
}

func exitErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func waitExit() {
	sigS := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigS, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigS
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("ccrontab start ...,and waiting some code")
	<-done
	fmt.Println("exiting")
}
