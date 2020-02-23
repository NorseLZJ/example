package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/NorseLZJ/example/get_code/config"
)

var (
	cfg = flag.String("conf", "./share.json", "share Server Config")
)

type FileSrc struct {
	ShareDir string `json:"share_dir"`
	Addr     string `json:"addr"`
}

func main() {
	addrS, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	addrTmpS := make([]string, 0)

	for _, v := range addrS {
		if ipNet, ok := v.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			addrTmpS = append(addrTmpS, ipNet.IP.String())
		}
	}

	fmt.Println("your computer ip")
	for _, v := range addrTmpS {
		fmt.Println(v)
	}

	fmt.Printf("\nstart server\n")
	cfgT := &FileSrc{}
	err = config.Marshal(*cfg, cfgT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfgT.Addr)
	fmt.Println(cfgT.ShareDir)
	log.Fatal(http.ListenAndServe(cfgT.Addr, http.FileServer(http.Dir(cfgT.ShareDir))))
}
