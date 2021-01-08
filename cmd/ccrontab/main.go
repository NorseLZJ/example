package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

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
	if err != nil {
		exitErr(err)
	}
	err = json.Unmarshal(data, cfgT)
	if err != nil {
		exitErr(err)
	}
	fd, err := os.OpenFile(*logDir, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		exitErr(err)
	}
	defer fd.Close()
	log.SetOutput(fd)
	c := cron.New(cron.WithSeconds())
	for _, vv := range cfgT.Tasks {
		_, err = c.AddFunc(vv.Space, func() {
			script, err := ioutil.ReadFile(vv.Script)
			if err != nil {
				exitErr(err)
			}
			cmd := exec.Command(vv.Cmd, vv.Param, string(script))
			b, err := cmd.Output()
			if err != nil {
				log.Error("cmd:%s script:%s Err: %v\n", vv.Cmd, vv.Script, err)
			}
			log.Info("\nCCB START TIME(%s)\n", time.Now().String())
			log.Info(fmt.Sprintf("%s", string(b)))
			log.Info("\nCCB END TIME(%s)\n", time.Now().String())
		})
		if err != nil {
			exitErr(err)
		}
	}
	c.AddFunc("*/5 * * * * ?", func() {
		log.Info("hello world")
	})
	c.Start()
	waitExit()
}

func exitErr(err error) {
	fmt.Println(err)
	os.Exit(1)
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
