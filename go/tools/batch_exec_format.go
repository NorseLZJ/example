package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	file   = flag.String("file", "", "file")
	dir    = flag.String("dir", "", "path")
	ignore = flag.String("skip", "", "skip file and dir")
)

var (
	ignorePatter = map[string]bool{
		"Debug":      true,
		".svn":       true,
		".git":       true,
		"build":      true,
		"contrib":    true,
		".vs":        true,
		".gitignore": true,
		".cd":        true,
		".rc":        true,
		".sln":       true,
		".db":        true,
		".filters":   true,
		".user":      true,
		".vcxproj":   true,
		".lib":       true,
	}
)

func main() {
	flag.Parse()

	if *ignore != "" {
		for _, i := range strings.Split(*ignore, ",") {
			ignorePatter[i] = true
		}
	}

	if *file != "" {
		//tocOde(*file, "gb2312", "utf-8")
		cmd := exec.Command("iconv", "-f GB2312", "-t UTF-8", "-o", *file, *file)
		if err := cmd.Start(); err != nil {
			fmt.Printf("(%s) -> Err : %v\n", *file, err)
		}
		if err := cmd.Wait(); err != nil {
			fmt.Printf("(%s) -> wait err : %v\n", *file, err)
		}
	} else {
		ret, err := GetAll(*dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, vv := range ret {
			cmd := exec.Command("iconv", "-f GB2312", "-t UTF-8", "-o", vv, vv)
			if err := cmd.Start(); err != nil {
				fmt.Printf("(%s) -> Err : %v\n", *file, err)
			}
			if err := cmd.Wait(); err != nil {
				fmt.Printf("(%s) -> wait err : %v\n", *file, err)
			}
		}
	}
}

func GetAll(path string) ([]string, error) {
	ret := make([]string, 0, 0)
	read, err := os.ReadDir(path)
	if err != nil {
		return ret, err
	}
	for _, fi := range read {
		if fi.IsDir() && !ignorePatter[fi.Name()] {
			fullDir := path + "/" + fi.Name()
			vvv, _ := GetAll(fullDir)
			ret = append(ret, vvv...)
		} else {
			if _, ok := ignorePatter[fi.Name()]; ok {
				continue
			}
			strs := strings.Split(fi.Name(), ".")
			if len(strs) < 2 {
				continue
			}
			sub := fmt.Sprintf(".%s", strs[1])
			if _, ok := ignorePatter[sub]; ok {
				continue
			}
			file := path + "/" + fi.Name()
			ret = append(ret, file)
		}
	}
	return ret, nil
}
