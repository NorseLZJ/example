package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
)

var (
	port = flag.String("p", "8900", "server port")
	help = flag.Bool("h", false, "help ")
)

const (
	shareWindows = "D:\\share"
	windows      = `windows`
	linux        = `linux`
	mac          = `darwin`
)

var cc = `
you need create share dir before run fsv.exe 

usage: fsv [-p]
Example:
1、fsv 
2、fsv -p 9000 

Default:
Port: 8900

ShareDir
Windows: 	"D:\share\"
Linux&MAC: 	"~/share/"
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
	str, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	curDirs := strings.Split(str, "/")
	user, err2 := user.Current()
	if err2 != nil {
		log.Fatal(err)
	}

	if len(curDirs) < 3 && runtime.GOOS != windows && user.Username != "root" {
		log.Fatal("please go to your home directory and run fsv")
	}
	switch runtime.GOOS {
	case windows:
		share = shareWindows
	case linux:
		if user.Username == "root" {
			share = "/root/share/"
		} else {
			share = "/home/" + curDirs[2] + "/share/"
		}
	case mac:
		share = "/Users/" + curDirs[2] + "/share/"
	default:
		log.Fatal("share dir is nil")
	}

	fmt.Println("start server")
	cpPort := fmt.Sprintf(":%s", *port)
	showIp(cpPort)
	err = http.ListenAndServe(cpPort, http.FileServer(http.Dir(share)))
	if err != nil {
		log.Fatal(err)
	}
}

func showIp(port string) {
	addrS, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("Get InterfaceAddrs err:%v", err))
	}
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
