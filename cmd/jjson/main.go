package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	org = flag.String("f", "", "file")
)

func main() {
	flag.Parse()
	if "" == *org {
		log.Fatal("file is nil")
	}
	b, err := ioutil.ReadFile(*org)
	if err != nil {
		log.Fatal("ReadAll err:", err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "")
	fd, err := os.OpenFile(*org, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal("Open err:", err)
	}
	defer fd.Close()
	out.WriteTo(fd)
}
