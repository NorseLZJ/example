package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log2 "github.com/NorseLZJ/example/std/log"
)

func main() {
	llog := log2.LLog()
	url := "https://sh.lianjia.com/ershoufang/"
	resp, err := http.Get(url)
	if err != nil {
		llog.Error("%v", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		llog.Error("%v", err)
	}

	if resp.StatusCode != 200 {
		llog.Error("status code err %d", resp.StatusCode)
	}

	if len(respBody) == 0 {
		respBody = []byte("{}")
	}

	var wResp interface{}
	if err = json.Unmarshal(respBody, &wResp); err != nil {
		llog.Error("Unmarshal err ", err)
	}
}
