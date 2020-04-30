package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	org = flag.String("f", "", "file")
	dir = flag.String("d", "", "Absolute path")
)

func main() {
	flag.Parse()
	if "" == *org && "" == *dir {
		log.Fatal("file is nil")
	}
	var err error
	if "" != *org {
		if err = jsonFormat(*org); err != nil {
			log.Fatal(err)
		}
	} else if "" != *dir {
		fs, err := ioutil.ReadDir(*dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range fs {
			if !f.IsDir() {
				ffdir := fmt.Sprintf("%s/%s", *dir, f.Name())
				err = jsonFormat(ffdir)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func jsonFormat(f string) error {
	b, err := ioutil.ReadFile(f)
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
