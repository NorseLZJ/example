package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/user"
	"runtime"
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

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	switch runtime.GOOS {
	case windows:
		share = shareWindows
	case linux:
		if user.Username == "root" {
			share = "/root/share/"
		} else {
			share = fmt.Sprintf("/home/%s/share", user.Username)
		}
	case mac:
		share = fmt.Sprintf("/Users/%s/share", user.Username)
	default:
		log.Fatal("share dir is nil")
	}

	fmt.Printf("share:%s\n", share)
	fmt.Println("start server")
	cpPort := fmt.Sprintf(":%s", *port)
	showIp(cpPort)

	fs := CustomFileServer(http.Dir(share))
	err = http.ListenAndServe(cpPort, RequestLogger(fs))
	if err != nil {
		fmt.Println(err)
	}
}
