package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	org = flag.String("f", "", "file")
	dir = flag.String("d", "", "Absolute path")
)

var usage = func() {
	printUsage(os.Stderr)
	os.Exit(1)
}

func printUsage(w *os.File) {
	fmt.Fprintf(w, "usage: jjson\n")
	fmt.Fprintf(w, "jjson -f xxx.json\n")
	fmt.Fprintf(w, "jjson -d xxx(dir)\n")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printUsage(os.Stdout)
		os.Exit(0)
	}
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		usage()
	}
	if "" == *org && "" == *dir {
		flag.Usage()
		log.Fatal("file and dir is nil ")
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
			if !f.IsDir() && strings.HasSuffix(f.Name(), ".json") {
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
