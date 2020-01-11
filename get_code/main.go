package main

import (
	"flag"
	"fmt"
	"log"
	"example/get_code/config"
)

var (
	cfg = flag.String("conf", "./goGet.conf", "go get file")
)

func main() {
	cfgT, err := config.Marshal(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range cfgT.Get {
		fmt.Println(v)
	}
	fmt.Println(cfgT.Proxy)
}
