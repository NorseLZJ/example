package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	file   = flag.String("f", "", "file")
	dir    = flag.String("d", "", "path")
	ignore = flag.String("i", "", "ignore file and dir")
	help   = flag.Bool("h", false, "help")
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

var usage = func() {
	printUsage(os.Stderr)
	os.Exit(1)
}

func printUsage(w *os.File) {
	info := `
vsfmt: format visual studio project file GB2312 -> utf8.
Usage :
vsfmt -f xxx.xxx(file);
or 
vsfmt -d xxx(dir);

some other param:
-h : print help info
-i : ignore file and dir patter ,like .svn,.git,.so, ### use "," split 
`
	fmt.Printf("%s\n", info)
	fmt.Println("default ignore")
	for vv, _ := range ignorePatter {
		fmt.Println(vv)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
	}

	if *file == "" && *dir == "" {
		flag.Usage()
	}

	if *ignore != "" {
		for _, i := range strings.Split(*ignore, ",") {
			ignorePatter[i] = true
		}
	}

	if *file != "" {
		//tocOde(*file, "gb2312", "utf-8")
		cmd := exec.Command("iconv", "-f GB2312", "-t UTF-8", "-o", *file, *file)
		err := cmd.Start()
		if err != nil {
			fmt.Printf("(%s) -> Err : %v\n", *file, err)
		}
		err = cmd.Wait()
		if err != nil {
			fmt.Printf("(%s) -> wait err : %v\n", *file, err)
		}
	} else {
		ret, err := GetAll(*dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, vv := range ret {
			cmd := exec.Command("iconv", "-f GB2312", "-t UTF-8", "-o", vv, vv)
			err := cmd.Start()
			if err != nil {
				fmt.Printf("(%s) -> Err : %v\n", vv, err)
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Printf("(%s) -> wait err : %v\n", vv, err)
			}
			//tocOde(vv, "gb2312", "utf-8")
		}
	}
}

func GetAll(path string) ([]string, error) {
	ret := make([]string, 0, 0)
	read, err := ioutil.ReadDir(path)
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

//func tocOde(f, formEncoding, toEnconding string) error {
//	b, err := ioutil.ReadFile(f)
//	if err != nil {
//		return err
//	}
//
//	out := make([]byte, len(b), len(b))
//	_, _, err = iconv.Convert(b, out, formEncoding, toEnconding)
//	if err != nil {
//		fmt.Println(err)
//		return err
//	}
//	//iconv.Convert(b, out, formEncoding, toEnconding)
//	fd, err := os.OpenFile(f, os.O_WRONLY|os.O_TRUNC, 0600)
//	if err != nil {
//		return err
//	}
//	defer fd.Close()
//
//	if _, err = fd.Write(out); err != nil {
//		return err
//	}
//	return nil
//}
//
