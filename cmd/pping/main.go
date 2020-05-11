package main

import (
	"flag"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/NorseLZJ/example/std/cfg_marshal"
)

var (
	sPing = flag.String("f", "./s_ping.json", "mping config")
)

func main() {
	flag.Parse()
	cfgT := &cfg_marshal.MPing{}
	err := cfg_marshal.Marshal(*sPing, cfgT)
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
