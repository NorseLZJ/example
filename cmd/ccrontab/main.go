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

	"github.com/robfig/cron"
)

type Task struct {
	Cmd    string `json:"cmd"`    // sh
	Param  string `json:"param"`  // -c
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
		exitErr(err)
	}
	err = json.Unmarshal(data, cfgT)
	if err != nil {
		exitErr(err)
	}

	c := cron.New(cron.WithSeconds())

	for _, vv := range cfgT.Tasks {
		_, err = c.AddFunc(vv.Space, func() {
			//cmd := exec.Command(vv.Cmd, vv.Param, vv.Script)
			cmd := exec.Command("sh", "-c", vv.Script)
			b, err := cmd.Output()
			if err != nil {
				log.Printf("cmd:%s script:%s Err: %v\n", vv.Cmd, vv.Script, err)
			}
			fmt.Println(string(b))
			//err = cmd.Start()
			//if err != nil {
			//	log.Printf("cmd:%s script:%s Err: %v\n", vv.Cmd, vv.Script, err)
			//}
			//err = cmd.Wait()
			//if err != nil {
			//	log.Printf("cmd:%s script:%s Err: %v\n", vv.Cmd, vv.Script, err)
			//}
		})
		if err != nil {
			exitErr(err)
		}
	}
	//_, err = c.AddFunc("*/5 * * * * ?", func() { fmt.Println("hello world") })
	//if err != nil {
	//	exitErr(err)
	//}
	c.Start()
	select {}
	//waitExit()
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
