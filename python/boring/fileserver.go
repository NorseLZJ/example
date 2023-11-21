package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	host = flag.String("host", ":8899", "listen host")
)

func main() {
	flag.Parse()
	err := http.ListenAndServe(*host, http.FileServer(http.Dir("D:\\share")))
	if err != nil {
		log.Fatal(err)
	}
}
