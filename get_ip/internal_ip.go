package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/NorseLZJ/example/cfg_marshal"
)

var (
	internalIp = flag.String("conf", "./internal_ip.json", "active ip config")
)

func main() {
	cfgT := &cfg_marshal.Active{}
	err := cfg_marshal.Marshal(*internalIp, cfgT)
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
