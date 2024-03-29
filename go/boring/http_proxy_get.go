package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

var (
	proxyUrl, _ = url.Parse("http://127.0.0.1:7890")
)

func main() {
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	client := &http.Client{Transport: transport}
	resp, err := client.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
