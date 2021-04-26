package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

var (
	dir    = flag.String("d", "./", "start folder")
	txtDir = flag.String("td", "Placements/", "txtDir")
	outDir = flag.String("od", "/out/", "outDir")
)

var (
	r = regexp.MustCompile(`"mc":.*"events":\[]}}}`)
)

func main() {
	flag.Parse()
	if *dir == "" || *txtDir == "" {
		panic("Must be fill dir and txtDir")
	}
	cwd, _ := os.Getwd()
	infos, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
	}
	saveDir := fmt.Sprintf("%s/%s", cwd, *outDir)
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err = os.Mkdir(saveDir, 0777); err != nil {
			panic(err)
		}
	}
	for _, txt := range infos {
		if !strings.HasSuffix(txt.Name(), ".json") {
			continue
		}
		data, err := ioutil.ReadFile(txt.Name())
		if err != nil {
			log.Printf("ReadFile file:[%s] err:[%s]\n", txt.Name(), err.Error())
			continue
		}
		str, orgStr, err := getStr(data)
		if err != nil {
			fmt.Printf("%s", err.Error())
			continue
		}
		ret := &ToolStruct{}
		if err = json.Unmarshal([]byte(orgStr), ret); err != nil {
			log.Printf("json.Unmarshal file:[%s] err:[%s]\n", txt.Name(), err.Error())
		}
		for _, res := range ret.Frames {
			tmpDir := fmt.Sprintf("%s%s%s.txt", *dir, *txtDir, res.Res)
			rData, err := ioutil.ReadFile(tmpDir)
			if err != nil {
				continue
			}
			sData := strings.Replace(string(rData), "\r\n", ";", -1)
			x, y, err := getTwoNum(sData)
			if err != nil {
				log.Printf("getTwoNum err file:[%s] txtFile[%s] err:[%s]\n", txt.Name(), res.Res, err.Error())
				continue
			}
			res.X = float32(x)
			res.Y = float32(y)
		}
		wData, err := json.Marshal(ret)
		if err != nil {
			continue
		}
		wStr := string(wData)
		wStr = strings.Replace(str, orgStr, wStr, -1)
		outFile := fmt.Sprintf("%s/%s/%s", cwd, *outDir, txt.Name())
		err = ioutil.WriteFile(outFile, []byte(wStr), 0644)
		if err != nil {
			log.Fatal(fmt.Sprintf("ioutil.WriteFile file:[%s] err:[%s]\n", outFile, err.Error()))
		}
		fmt.Printf("Write Success [%s]\n", txt.Name())
	}
	waitExit()
}

func getStr(src []byte) (string, string, error) {
	str := string(src)
	str = strings.Replace(str, "\r\n", "", -1)
	str = strings.Replace(str, " ", "", -1)

	v := r.FindAllString(str, -1)
	if len(v) != 1 {
		return "", "", errors.New(fmt.Sprintf("Regexp err [%s]\n", string(src)))
	}
	newStr := v[0]
	newStr = strings.Replace(newStr, "mc\":{", "", -1)
	newStr = newStr[strings.Index(newStr, "{") : len(newStr)-2]
	return str, newStr, nil
}

func getTwoNum(val string) (int, int, error) {
	strS := strings.Split(val, ";")
	if len(strS) < 2 {
		return 0, 0, errors.New("can't find match charset")
	}

	v1, err1 := strconv.Atoi(strS[0])
	if err1 != nil {
		return 0, 0, err1
	}
	v2, err2 := strconv.Atoi(strS[1])
	if err1 != nil {
		return 0, 0, err2
	}
	return v1, v2, nil
}

func waitExit() {
	sigS := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigS, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigS
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("处理完了，可以关掉程序 ctrl+c 或 点击关闭按钮（右上角的X）")
	<-done
	fmt.Println("exiting")
}
