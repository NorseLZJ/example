package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/NorseLZJ/example/get_code/config"
)

var (
	cfg = flag.String("conf", "./active.json", "active ip config")
)

type Active struct {
	Min      int    `json:"Min"`
	Max      int    `json:"Max"`
	BaseAddr string `json:"BaseAddr"`
}

func main() {
	cfgT := &Active{}
	err := config.Marshal(*cfg, cfgT)
	if err != nil {
		log.Fatal(err)
	}
	ipList := make([]string, 0, cfgT.Max-cfgT.Min)
	for i := cfgT.Min; i < cfgT.Max; i++ {
		ipList = append(ipList, fmt.Sprintf("%s.%d", cfgT.BaseAddr, i))
	}

	activeIp := make([]string, 0)

	wg := sync.WaitGroup{}
	wg.Add(cfgT.Max - cfgT.Min)
	for _, v := range ipList {
		go func(ip string) {
			defer wg.Done()
			cmd := exec.Command("ping", "-c", "1", ip)
			err := cmd.Start()
			if err != nil {
				//log.Println("Start cmd ", err)
				return
			}
			err = cmd.Wait()
			if err != nil {
				//log.Println("Wait cmd ", err)
				return
			}
			activeIp = append(activeIp, ip)
		}(v)
	}

	wg.Wait()

	fmt.Println("ActiveIp -------")
	for _, ip := range activeIp {
		fmt.Println(ip)
	}
}
