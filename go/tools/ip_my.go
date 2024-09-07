package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"sync"
)

const (
	min = 1
	max = 256
)

func main() {
	flag.Parse()
	activeIp := make(map[int][]string, 0)
	wg := sync.WaitGroup{}
	ipList := internalIp()
	maybeIp := maybeActiveIp(ipList)
	for idx, ippList := range maybeIp {
		activeIp[idx] = []string{}
		for _, ip := range ippList {
			wg.Add(1)
			go func(ip string, idx int) {
				defer wg.Done()
				cmd := exec.Command("ping", "-c", "1", ip)
				err := cmd.Start()
				if err != nil {
					return
				}
				err = cmd.Wait()
				if err != nil {
					return
				}
				activeIp[idx] = append(activeIp[idx], ip)
			}(ip, idx)
		}
	}
	fmt.Println("wiat .......")
	wg.Wait()
	fmt.Println("Internal ActiveIp")
	for idx, ll := range activeIp {
		if len(ll) <= 0 {
			continue
		}
		fmt.Printf("ipGroup:%d\n", idx)
		for _, ip := range ll {
			fmt.Printf("%s\n", ip)
		}
		fmt.Println()
	}
}

func maybeActiveIp(bIp []string) map[int][]string {
	ret := make(map[int][]string)
	for idx, ipp := range bIp {
		ret[idx] = []string{}
		for i := min; i < max; i++ {
			vv := fmt.Sprintf("%s.%d", ipp, i)
			ret[idx] = append(ret[idx], vv)
		}
	}
	return ret
}

func internalIp() []string {
	addrS, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	addrTmpS := make([]string, 0)
	for _, v := range addrS {
		if ipNet, ok := v.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			cpIp := ipNet.IP.String()
			ipSlice := strings.Split(cpIp, ".")[:3]
			aip := fmt.Sprintf("%s.%s.%s", ipSlice[0], ipSlice[1], ipSlice[2])
			addrTmpS = append(addrTmpS, aip)
		}
	}
	return addrTmpS
}
