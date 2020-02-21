package main

import (
	"flag"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/NorseLZJ/example/get_code/config"
)

var (
	cfg = flag.String("conf", "./mping.json", "mping config")
)

type MPing struct {
	Frequency int    `json:"Frequency"`
	Sleep     int    `json:"Sleep"`
	Size      string `json:"Size"`
	Addr      string `json:"Addr"`
}

func main() {
	cfgT := &MPing{}
	err := config.Marshal(*cfg, cfgT)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(cfgT.Frequency)
	for i := 0; i < cfgT.Frequency; i++ {
		time.Sleep(time.Microsecond * 500)
		go func() {
			defer wg.Done()
			cmd := exec.Command("ping", "-c", "5", "-s", cfgT.Size, cfgT.Addr)
			err := cmd.Start()
			if err != nil {
				log.Println("Start cmd ", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				log.Println("Wait cmd ", err)
				return
			}
		}()
	}

	wg.Wait()
}
