package main

import (
	"flag"
	"fmt"
	"log"
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
	cfgT := &FileSrc{}
	err := config.Marshal(*cfg, cfgT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfgT.Addr)
	fmt.Println(cfgT.ShareDir)
	log.Fatal(http.ListenAndServe(cfgT.Addr, http.FileServer(http.Dir(cfgT.ShareDir))))
}
