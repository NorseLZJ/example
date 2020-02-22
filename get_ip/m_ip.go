package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addrS, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	addrTmpS := make([]string, 0, 0)

	for _, v := range addrS {
		if ipNet, ok := v.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() &&
			ipNet.IP.To4() != nil {

			addrTmpS = append(addrTmpS, ipNet.IP.String())
		}
	}

	for _, v := range addrTmpS {
		fmt.Println(v)
	}
}
