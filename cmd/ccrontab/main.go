package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
)

type Task struct {
	Cmd    string `json:"cmd"`    // bin/bash
	Script string `json:"script"` // /aaa/bbb.sh
	Space  string `json:"space"`  // * * * * * *
}

type Config struct {
	LogDir string `json:"logDir"`
	Tasks  []Task `json:"tasks"`
}

var (
	cfg = flag.String("c", "./ccrontab.json", "crontab config json")
)

func main() {
	flag.Parse()
	cfgT := &Config{}
	data, err := ioutil.ReadFile(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, cfgT)
	if err != nil {
		log.Fatal(err)
	}

	c := cron.New()
	c.AddFunc("* * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.Start()
	waitExit()
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
