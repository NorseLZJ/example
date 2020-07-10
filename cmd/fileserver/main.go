package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/NorseLZJ/example/std"
)

var (
	port = flag.String("p", "8900", "server port")
	help = flag.Bool("h", false, "help ")
)

const (
	shareWindows = "D:\\share"
	windows      = `windows`
	linux        = `linux`
)

var cc = `
you need create share dir before run fsv.exe 

usage: fsv [-p]
Example:
1、fsv 
2、fsv -p 9000 

Default:
Windows: share dir 	"D:\share\"
Linux: share dir 	"~/share/"
`

func main() {
	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("%s\n", cc)
	}
	if *help {
		flag.Usage()
		return
	}
	share := ""
	switch runtime.GOOS {
	case windows:
		share = shareWindows
	case linux:
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		strs := strings.Split(path, "/")
		cpStr := strings.Join(strs[0:3], "/")
		cpStr += "/share"
		share = cpStr
	default:
		log.Fatal("share dir is nil")
	}

	fmt.Println("start server")
	cpPort := fmt.Sprintf(":%s", *port)
	showIp(cpPort)
	err := http.ListenAndServe(cpPort, http.FileServer(http.Dir(share)))
	if err != nil {
		log.Fatal(err)
	}
}

func showIp(port string) {
	addrS, err := net.InterfaceAddrs()
	std.CheckErr(err)
	addrTmpS := make([]string, 0)
	for _, v := range addrS {
		if ipNet, ok := v.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			addrTmpS = append(addrTmpS, ipNet.IP.String())
		}
	}
	fmt.Println("try access this address please")
	for _, v := range addrTmpS {
		fmt.Println(fmt.Sprintf("%s%s", v, port))
	}
}
