package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	org = flag.String("file", "", "file")
	dir = flag.String("dir", "", "dir")
)

func main() {
	flag.Parse()
	if *org != "" {
		if err := jsonFormat(*org); err != nil {
			log.Fatal(err)
		}
	} else if *dir != "" {
		fs, err := os.ReadDir(*dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range fs {
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
				f := fmt.Sprintf("%s/%s", *dir, f.Name())
				if err := jsonFormat(f); err != nil {
					log.Printf("file:%s err:%s\n", f, err.Error())
				}
			}
		}
	} else {
		flag.Usage()
	}
}

func jsonFormat(f string) error {
	b, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err = json.Indent(&out, b, "", ""); err != nil {
		return err
	}
	fd, err := os.OpenFile(f, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer fd.Close()
	if _, err = out.WriteTo(fd); err != nil {
		return err
	}
	return nil
}
