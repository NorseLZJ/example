package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("netcmd.proto")
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`message(.*?)}`)
	rfe := regexp.MustCompile(`[usir](.*?);`)
	strData := string(data)
	strData = strings.ReplaceAll(strData, "\n", "")
	strData = strings.ReplaceAll(strData, "\r\n", "")
	strData = strings.ReplaceAll(strData, "  ", " ")
	strData = strings.ReplaceAll(strData, "   ", " ")
	writeFile := ""
	rets := r.FindAllString(strData, -1)
	for _, v := range rets {

		strMsg := strings.ReplaceAll(v, "message", "")

		idx1 := strings.Index(strMsg, "{")
		idx2 := strings.Index(strMsg, "}")
		if idx2-idx1 < 2 {
			continue
		}
		msgName := strings.ReplaceAll(strMsg[0:idx1], " ", "")
		strMsg = strMsg[idx1:idx2]

		frets := rfe.FindAllString(strMsg, -1)

		outText := fmt.Sprintf("void TEST_%s() \n{\n\tNetCmd::%s info;", msgName, msgName)
		for idx, field := range frets {

			field = strings.ReplaceAll(field, "\t", " ")
			field = strings.ReplaceAll(field, "  ", " ")
			field = strings.ReplaceAll(field, "   ", " ")
			vvs := strings.Split(field, " ")
			if len(vvs) <= 2 {
				continue
			}
			if vvs[0] == "repeated" && len(vvs) >= 3 {
				setName := strings.ToLower(vvs[2])
				if idx == 0 {
					outText = fmt.Sprintf("%s\n\tinfo.set_allocated_%s();", outText, setName)
				} else {
					outText = fmt.Sprintf("%s\n\tinfo.set_allocated_%s();", outText, setName)
				}
			} else {
				setName := strings.ToLower(vvs[1])
				if idx == 0 {
					outText = fmt.Sprintf("%s\n\tinfo.set_%s();", outText, setName)
				} else {
					outText = fmt.Sprintf("%s\n\tinfo.set_%s();", outText, setName)
				}
			}
		}
		outText = fmt.Sprintf("\n%s\n}", outText)
		writeFile += outText
	}
	os.WriteFile("Test_NetCmd.cc", []byte(writeFile), 0777)
}
