package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	dir    = flag.String("d", "./", "start folder")
	txtDir = flag.String("td", "Placements/", "txtDir")
)

var (
	r = regexp.MustCompile(`"mc":{"":.*"events":\[]}}}`)
)

func main() {
	flag.Parse()
	if *dir == "" || *txtDir == "" {
		panic("Must be fill dir and txtDir")
	}
	infos, err := ioutil.ReadDir(*dir)
	if err != nil {
		panic(err)
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
		str := string(data)
		str = strings.Replace(str, "\r\n", "", -1)
		str = strings.Replace(str, " ", "", -1)

		v := r.FindAllString(str, -1)
		if len(v) != 1 {
			log.Printf("Regexp err [%s]\n", string(data))
			continue
		}
		newStr := v[0]
		newStr = strings.Replace(newStr, "\"mc\":{\"\":", "", -1)
		newStr = newStr[0 : len(newStr)-2]
		orgStr := newStr // TODO replace after
		ret := &ToolStruct{}
		if err = json.Unmarshal([]byte(newStr), ret); err != nil {
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
				continue
			}
			res.X = x
			res.Y = y
		}
		wData, err := json.Marshal(ret)
		if err != nil {
			continue
		}
		wStr := string(wData)
		wStr = strings.Replace(str, orgStr, wStr, -1)
		err = ioutil.WriteFile(txt.Name(), []byte(wStr), 0644)
		if err != nil {
			log.Printf("ioutil.WriteFile file:[%s] err:[%s]\n", txt.Name(), err.Error())
		}
		fmt.Printf("Write Success [%s]\n", txt.Name())
	}
}

func getTwoNum(val string) (int, int, error) {
	strS := strings.Split(val, ";")
	if len(strS) != 2 {
		return 0, 0, errors.New("can't match!")
	}

	v1, err1 := strconv.Atoi(strS[0])
	if err1 != nil {
		return 0, 0, err1
	}
	v2, err2 := strconv.Atoi(strS[0])
	if err1 != nil {
		return 0, 0, err2
	}
	return v1, v2, nil
}
